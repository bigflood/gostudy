package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/bigflood/gostudy/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":8888"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	hostname, _ := os.Hostname()

	s := grpc.NewServer()
	pb.RegisterPingServer(s, &server{hostname: hostname})
	reflection.Register(s)

	log.Println(hostname, " serving: ", addr, " ..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	hostname string
}

func (s *server) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingRes, error) {
	res := &pb.PingRes{
		Msg: s.hostname + ": " + req.Msg + " pong",
	}

	return res, nil
}
