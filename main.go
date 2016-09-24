package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker World!\n\nYour IP: %s", r.RemoteAddr)
}

func main() {
	listenAddr := "0.0.0.0:8080"

	log.Println("Starting HTTP server on", listenAddr)
	http.HandleFunc("/", handler)
	http.ListenAndServe(listenAddr, nil)
}
