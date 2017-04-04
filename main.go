package main

import (
	"log"
	"net/http"
)

const (
	Port = ":9000"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		NotImplemented(w)
	})

	http.HandleFunc("/registration", HandleRegistration)

	log.Printf("Starting server on localhost%s", Port)
	http.ListenAndServe(Port, nil)
}
