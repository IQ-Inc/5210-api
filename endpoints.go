package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// NotImplemented returns a not implemented status to the client
// Many things are not implemented...
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// HandleRegistration handle user registration, storing data in the database.
func HandleRegistration(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		// Parse registration payload
		decoder := json.NewDecoder(r.Body)

		var reg Registration

		if err := decoder.Decode(&reg); nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Failed to decode JSON in HandleRegistration")
			return
		} else if reg.ID == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Printf("Registration requires UUID")
			return
		}

		reg.table = "Registration"
		reg.timestamp = time.Now()

		err := Store(ctx, &reg)
		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprint(w, "OK")

	// No other HTTP verbs implemented
	default:
		NotImplemented(w, r)
	}
}
