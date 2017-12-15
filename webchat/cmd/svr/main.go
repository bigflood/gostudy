package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bigflood/gostudy/webchat/svr"
)

func main() {
	redisEndpoint := os.Getenv("REDIS_ENDPOINT")

	svr := svr.New(redisEndpoint)

	http.HandleFunc("/webchat", svr.WebchatHandler)

	dn, _ := filepath.Abs("static")
	log.Println("static: ", dn)
	//http.Handle("/static", http.FileServer(http.Dir(dn)))
	http.Handle("/static/", http.StripPrefix("/static/", disableCacheHandler(http.FileServer(http.Dir(dn)))))

	//http.Handle("/", http.RedirectHandler("/static/index.html", 301))

	addr := ":8080"
	log.Println("listen:", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func disableCacheHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Add("Pragma", "no-cache")
		w.Header().Add("Expires", "0")
		h.ServeHTTP(w, r)
	})
}
