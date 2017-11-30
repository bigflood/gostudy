package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bigflood/gostudy/webchat/svr"
)

func main() {
	redisEndpoint := os.Getenv("REDIS_ENDPOINT")

	svr := svr.New(redisEndpoint)

	http.HandleFunc("/webchat", svr.WebchatHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
