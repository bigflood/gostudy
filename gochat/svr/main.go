package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"

	"github.com/go-redis/redis"
)

const chattingChan = "ch1"

func connectToRedis() *redis.Client {
	redisEndpoint := os.Getenv("REDIS_ENDPOINT")
	log.Println("redisEndpoint:", redisEndpoint)

	client := redis.NewClient(&redis.Options{
		Addr:     redisEndpoint,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

func main() {
	redisClient := connectToRedis()

	go http.ListenAndServe(":9090", nil)

	port := "8080"
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	log.Println("Listen", port, "..")

	svr := NewChattingServer(redisClient)

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
	redisClient *redis.Client

	lock    sync.Mutex
	clients []*clientInfo
}

type clientInfo struct {
	conn   net.Conn
	pubSub *redis.PubSub
}

func NewChattingServer(redisClient *redis.Client) *chattingServer {
	return &chattingServer{
		redisClient: redisClient,
	}
}

func (svr *chattingServer) AddClient(conn net.Conn) {
	info := &clientInfo{
		conn: conn,
	}

	info.pubSub = svr.redisClient.Subscribe(chattingChan)

	svr.lock.Lock()
	defer svr.lock.Unlock()
	svr.clients = append(svr.clients, info)

	go svr.recvFromClient(info)
	go svr.sendToClient(info)
}

func (svr *chattingServer) recvFromClient(info *clientInfo) {
	defer info.pubSub.Close()

	buf := make([]byte, 1024)

	for {
		n, err := info.conn.Read(buf)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		msg := string(buf[:n])
		_, err = svr.redisClient.Publish(chattingChan, msg).Result()
		if err != nil {
			log.Println("Publish error:", err)
		}
	}

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
	for {
		msg, err := info.pubSub.ReceiveMessage()
		if err != nil {
			log.Println("Receive error:", err)
			break
		}

		s := msg.Payload

		if _, err := info.conn.Write(([]byte)(s)); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
