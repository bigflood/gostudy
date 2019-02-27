package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	log.Print("123123", "a", "b", " \n", "sfd ")
	log.Println("123123", "a", "b"," \n", "sfd ")
	log.Printf("123123 %v", 23434)
	log.Fatal("234234")
}

