package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {

	svr := &server{}

	http.HandleFunc("/webchat", svr.webchatHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type server struct {
	clients []*client
}

type client struct {
}

func (svr *server) webchatHandler(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer c.Close()

	ch := make(chan []byte)
	defer close(ch)

	go func() {
		for msg := range ch {
			c.WriteMessage(websocket.BinaryMessage, msg)
		}
	}()
	// go svr.SendMessagesToClient(c)

mainLoop:
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		switch mt {
		case websocket.BinaryMessage:
			// svr.onRecvMsg(ch, msg)
			ch <- msg
		case websocket.CloseMessage:
			break mainLoop
		}
	}
}

func (svr *server) onRecvMsg(msg []byte) {
	// send msg to room
}

func (svr *server) SendMessagesToClient(c *websocket.Conn) {
	// send msg to room
}
