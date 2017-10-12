package main

import (
	"log"
	"net/http"

	"github.com/bigflood/gostudy/gogen/httptransport"
	"github.com/bigflood/gostudy/gogen/service"
)

func main() {
	svc := service.New()
	ht := httptransport.New(svc)
	log.Println("Listen..")
	http.ListenAndServe(":8080", ht)
}
