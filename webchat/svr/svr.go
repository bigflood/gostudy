package svr

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

type ChattingServer interface {
	WebchatHandler(w http.ResponseWriter, r *http.Request)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

const chattingChan = "ch1"

type chattingServer struct {
	hostname    string
	redisClient *redis.Client

	lock    sync.Mutex
	clients []*clientInfo
}

type clientInfo struct {
	conn   *websocket.Conn
	pubSub *redis.PubSub
}

func New(redisEndpoint string) ChattingServer {
	redisClient := connectToRedis(redisEndpoint)
	hostname, _ := os.Hostname()

	return &chattingServer{
		hostname:    hostname,
		redisClient: redisClient,
	}
}

func (svr *chattingServer) WebchatHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("WebchatHandler", r.URL)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade failed!", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	svr.AddClient(c)
}

func (svr *chattingServer) AddClient(conn *websocket.Conn) {
	log.Println("AddClient..")
	defer log.Println("AddClient.. end")

	info := &clientInfo{
		conn: conn,
	}

	info.pubSub = svr.redisClient.Subscribe(chattingChan)

	svr.lock.Lock()
	svr.clients = append(svr.clients, info)
	svr.lock.Unlock()

	info.conn.WriteMessage(websocket.TextMessage, ([]byte)(fmt.Sprint("hostname:", svr.hostname)))

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

		log.Println("recv", mt, msg)

		switch mt {
		case websocket.TextMessage, websocket.BinaryMessage:
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
		s = fmt.Sprint("from:", svr.hostname, "msg:", s)

		log.Println("WriteMessage..", s)

		if err := info.conn.WriteMessage(websocket.TextMessage, ([]byte)(s)); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func connectToRedis(redisEndpoint string) *redis.Client {
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
