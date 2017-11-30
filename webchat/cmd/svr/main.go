package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bigflood/gostudy/webchat/svr"
	"path/filepath"
)

func main() {
	redisEndpoint := os.Getenv("REDIS_ENDPOINT")

	svr := svr.New(redisEndpoint)


	http.HandleFunc("/webchat", svr.WebchatHandler)

	dn, _ := filepath.Abs("static")
	log.Println("static: ", dn)
	//http.Handle("/static", http.FileServer(http.Dir(dn)))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dn))))

	//http.Handle("/", http.RedirectHandler("/static/index.html", 301))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
