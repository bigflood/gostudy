package main

import (
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	go http.ListenAndServe(":9090", nil)

	port := "8080"
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	log.Println("Listen", port, "..")

	svr := NewChattingServer()
	go svr.RunMsgDistributer()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		log.Println("Accepted:", conn.RemoteAddr())

		svr.AddClient(conn)
	}
}

type chattingServer struct {
	lock    sync.Mutex
	clients []*clientInfo
	msgChan chan []byte
}

type clientInfo struct {
	conn     net.Conn
	sendChan chan []byte
}

func NewChattingServer() *chattingServer {
	return &chattingServer{
		msgChan: make(chan []byte),
	}
}

func (svr *chattingServer) RunMsgDistributer() {
	for msg := range svr.msgChan {

		svr.lock.Lock()
		for _, client := range svr.clients {
			client.sendChan <- msg
		}
		svr.lock.Unlock()
	}
}

func (svr *chattingServer) AddClient(conn net.Conn) {
	info := &clientInfo{
		conn:     conn,
		sendChan: make(chan []byte),
	}

	svr.lock.Lock()
	defer svr.lock.Unlock()
	svr.clients = append(svr.clients, info)

	go svr.recvFromClient(info)
	go svr.sendToClient(info)
}

func (svr *chattingServer) recvFromClient(info *clientInfo) {
	buf := make([]byte, 1024)

	for {
		n, err := info.conn.Read(buf)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		svr.msgChan <- buf[:n]
	}

	close(info.sendChan)
	info.conn.Close()

	svr.lock.Lock()
	defer svr.lock.Unlock()
	for i, p := range svr.clients {
		if p == info {
			svr.clients = append(svr.clients[:i], svr.clients[i+1:]...)
			break
		}
	}
}

func (svr *chattingServer) sendToClient(info *clientInfo) {
	for msg := range info.sendChan {
		_, err := info.conn.Write(msg)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
