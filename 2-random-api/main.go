package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handleRequest)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	randValue := rand.IntN(6) + 1
	w.Write([]byte(string(rune(randValue))))
	return
}
