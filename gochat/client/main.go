package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	svrAddr := flag.String("h", "localhost:8080", "server address")
	flag.Parse()

	log.Println("Connect to", *svrAddr, "..")
	conn, err := net.Dial("tcp", *svrAddr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go io.Copy(os.Stdout, conn)

	io.Copy(conn, os.Stdin)
}
