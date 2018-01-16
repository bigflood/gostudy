package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	addr := ":8888"
	router := mux.NewRouter()
	router.HandleFunc("/ping/{msg}", GetPing).Methods("GET")
	router.HandleFunc("/finger/{msg}", GetFinger).Methods("GET")

	log.Println("listen", addr, "..")
	log.Fatal(http.ListenAndServe(addr, router))

}

type Ping struct {
	Msg string `json:"msg,omitempty"`
}

type Finger struct {
	Msg string `json:"msg,omitempty"`
}

func GetPing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := Ping{Msg: params["msg"]}
	json.NewEncoder(w).Encode(data)
}

func GetFinger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := Finger{Msg: params["msg"]}
	json.NewEncoder(w).Encode(data)
}
