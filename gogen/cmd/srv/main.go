package main

import (
	"dadev/gostudy/gogen/httptransport"
	"dadev/gostudy/gogen/service"
	"log"
	"net/http"
)

func main() {
	svc := service.New()
	ht := httptransport.New(svc)
	log.Println("Listen..")
	http.ListenAndServe(":8080", ht)
}
