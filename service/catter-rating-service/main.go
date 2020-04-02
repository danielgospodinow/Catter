package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting up the server")

	registerHandlers()
	startServer()
}

func registerHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Root")
	})
}

func startServer() {
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
