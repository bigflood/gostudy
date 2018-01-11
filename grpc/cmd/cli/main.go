package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/bigflood/gostudy/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("addr", "localhost:8082", "endpoint")
	flag.Parse()

	log.Println("dial ", *addr, " ..")
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPingClient(conn)

	N := 20
	tsum := time.Duration(0)

	for i := 0; i < N; i++ {
		msg := fmt.Sprint("msg", i)

		t1 := time.Now()
		r, err := c.Ping(context.Background(), &pb.PingReq{Msg: msg})
		t2 := time.Now()
		d := t2.Sub(t1)
		tsum += d

		if err != nil {
			log.Fatalf("could not ping: %v", err)
		}

		log.Printf("Reply: %s (%s)", r.Msg, d)
	}

	log.Printf("avg time: %s", tsum/time.Duration(N))
}
