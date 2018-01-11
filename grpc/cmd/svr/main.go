package main

import (
	"context"
	"log"
	"net"

	"github.com/bigflood/gostudy/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":8082"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPingServer(s, &server{})
	reflection.Register(s)

	log.Println("Serving: ", addr, " ..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingRes, error) {
	res := &pb.PingRes{
		Msg: req.Msg + " pong",
	}

	return res, nil
}
