package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func NotImplemented(w http.ResponseWriter) {
	http.Error(w, "Not Implemented", http.StatusUnauthorized)
}

func HandleRegistration(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received request /registration -- Method: %s  Body: %s", r.Method, r.Body)

	switch r.Method {
	case http.MethodPost:

		// Parse registration payload
		decoder := json.NewDecoder(r.Body)

		var reg Registration

		if err := decoder.Decode(&reg); nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Failed to decode JSON in HandleRegistration")
			return
		} else if reg.UUID == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Printf("Registration requires UUID")
			return
		}

		err := StoreRequest(reg)
		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprint(w, "OK")

	default:
		NotImplemented(w)
	}
}
