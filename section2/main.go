package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	log.Println("start server")

	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("received request!")
	fmt.Fprintf(w, "Hello Docker!")
}