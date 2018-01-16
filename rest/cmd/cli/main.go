package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GoRest interface {
	Ping(msg string) (string, error)
}

func main() {
	addr := flag.String("addr", "localhost:8888", "endpoint")
	flag.Parse()

	log.Println("addr: ", *addr)
	c, err := newGoRest(*addr)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	N := 20
	tsum := 0.0
	tmin := 999.0
	tmax := 0.0
	count := 0

	for i := 0; i < N; i++ {
		msg := fmt.Sprint("msg", i)

		t1 := time.Now()
		r, err := c.Ping(msg)
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

		log.Printf("Reply: %s (%.3f ms)", r, d)
	}

	log.Printf("ping x %d : avg time=%.6f min=%.6f max=%.6f ms", count, tsum/float64(count), tmin, tmax)
}

func newGoRest(addr string) (GoRest, error) {
	return &client{addr: addr}, nil
}

type client struct {
	addr string
}

type Ping struct {
	Msg string `json:"msg,omitempty"`
}

func (c *client) Ping(msg string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/ping/%s", c.addr, msg))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	defer resp.Body.Close()

	reply := Ping{}

	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return "", err
	}
	return reply.Msg, nil
}
