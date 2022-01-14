package main

import (
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("ping")
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
