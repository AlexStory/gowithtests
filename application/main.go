package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}