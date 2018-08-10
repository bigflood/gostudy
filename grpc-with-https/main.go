package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/bigflood/gostudy/gencert"
	"github.com/bigflood/gostudy/grpc-with-https/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//go:generate protoc ./pb/ping.proto --go_out=plugins=grpc:.

func main() {
	//grpc.EnableTracing = true

	serveWithTLS()

	//http.ListenAndServe(":9999", nil)
}

func serveWithTLS() {
	cert, err := gencert.Generate("localhost", false)
	if err != nil {
		panic(err)
	}

	addr := startServer(cert)

	certPool := x509.NewCertPool()
	addCert(certPool, cert.CertBytes)

	clientCreds := credentials.NewTLS(
		&tls.Config{
			RootCAs:            certPool,
			InsecureSkipVerify: false,
		},
	)

	if err := startClient(addr, grpc.WithTransportCredentials(clientCreds)); err != nil {
		log.Println(err)
	}
}

func addCert(s *x509.CertPool, pemCerts []byte) {
	for len(pemCerts) > 0 {
		var block *pem.Block
		block, pemCerts = pem.Decode(pemCerts)
		if block == nil {
			break
		}
		if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
			continue
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			continue
		}

		s.AddCert(cert)
	}
}

func startClient(addr string, opts ...grpc.DialOption) error {
	log.Println("dial.. ", addr)

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return err
	}

	defer conn.Close()

	log.Println("conn state: ", conn.GetState())

	client := pb.NewPingClient(conn)

	req := &pb.PingReq{Msg: "123123"}
	log.Println("ping.. ", req.Msg)
	res, err := client.Ping(context.Background(), req)
	if err != nil {
		return err
	}

	log.Println("res: ", res.Msg)
	return nil
}

func startServer(certData gencert.Data, opts ...grpc.ServerOption) string {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	grpcSvr := grpc.NewServer(opts...)
	pb.RegisterPingServer(grpcSvr, new(pingSvr))

	httpSvr := http.Server{
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{certData.Cert},
		},
		Handler: grpcSvr,
	}

	go func() {
		log.Println("serve..")
		httpSvr.ServeTLS(listener, "", "")
	}()

	return "localhost" + portOfAddr(listener.Addr())
}

func portOfAddr(a net.Addr) string {
	s := a.String()
	i := strings.LastIndex(s, ":")
	if i < 0 {
		return ""
	}

	return s[i:]
}

type pingSvr struct {
}

func (*pingSvr) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingRes, error) {
	return &pb.PingRes{Msg: "pong " + req.Msg}, nil
}
