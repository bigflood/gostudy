package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/go-redis/redis"
	"os"
	"fmt"
	"sync"
)

var upgrader = websocket.Upgrader{}
const chattingChan = "ch1"

func main() {
	redisClient := connectToRedis()

	svr := NewChattingServer(redisClient)

	http.HandleFunc("/webchat", svr.webchatHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type chattingServer struct {
	redisClient *redis.Client

	lock    sync.Mutex
	clients []*clientInfo
}

type clientInfo struct {
	conn   *websocket.Conn
	pubSub *redis.PubSub
}

func NewChattingServer(redisClient *redis.Client) *chattingServer {
	return &chattingServer{
		redisClient: redisClient,
	}
}

func (svr *chattingServer) webchatHandler(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	svr.AddClient(c)
}

func (svr *chattingServer) AddClient(conn *websocket.Conn) {
	info := &clientInfo{
		conn: conn,
	}

	info.pubSub = svr.redisClient.Subscribe(chattingChan)

	svr.lock.Lock()
	svr.clients = append(svr.clients, info)
	svr.lock.Unlock()

	go svr.sendToClient(info)

	svr.recvFromClient(info)
}

func (svr *chattingServer) recvFromClient(info *clientInfo) {
	defer info.pubSub.Close()
	defer info.conn.Close()

mainLoop:
	for {
		mt, msg, err := info.conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		switch mt {
		case websocket.BinaryMessage:
			_, err = svr.redisClient.Publish(chattingChan, msg).Result()
			if err != nil {
				log.Println("Publish error:", err)
			}
		case websocket.CloseMessage:
			break mainLoop
		}
	}

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

		if err := info.conn.WriteMessage(websocket.BinaryMessage, ([]byte)(s)); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}


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
