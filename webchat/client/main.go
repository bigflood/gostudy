package main

import (
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/webchat", nil)
	if err != nil {
		panic(err)
	}

	defer c.Close()

	buf := make([]byte, 256)

mainLoop:
	for {
		r, err := os.Stdin.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}

		msg := buf[:r]

		if err := c.WriteMessage(websocket.BinaryMessage, msg); err != nil {
			log.Println(err)
			break
		}

		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		switch mt {
		case websocket.BinaryMessage:
			os.Stdout.Write(msg)
		case websocket.CloseMessage:
			break mainLoop
		}
	}
}
