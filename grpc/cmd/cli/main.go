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
	addr := flag.String("addr", "localhost:8888", "endpoint")
	flag.Parse()

	log.Println("dial ", *addr, " ..")
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPingClient(conn)

	N := 20
	tsum := 0.0
	tmin := 999.0
	tmax := 0.0
	count := 0

	for i := 0; i < N; i++ {
		msg := fmt.Sprint("msg", i)

		t1 := time.Now()
		r, err := c.Ping(context.Background(), &pb.PingReq{Msg: msg})
		t2 := time.Now()
		d := float64(t2.Sub(t1)) / float64(time.Millisecond)

		if i > 0 {
			count++
			tsum += d
			if d < tmin {
				tmin = d
			}
			if d > tmax {
				tmax = d
			}
		}

		if err != nil {
			log.Fatalf("could not ping: %v", err)
		}

		log.Printf("Reply: %s (%.3f ms)", r.Msg, d)
	}

	log.Printf("ping x %d : avg time=%.6f min=%.6f max=%.6f ms", count, tsum/float64(count), tmin, tmax)
}
