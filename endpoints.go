package main

import (
	"encoding/json"
	"net/http"
)

func NotImplemented(w http.ResponseWriter) {
	http.Error(w, "Not Implemented", http.StatusUnauthorized)
}

func HandleRegistration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		// Parse registration payload
		decoder := json.NewDecoder(r.Body)

		var reg Registration

		if err := decoder.Decode(&reg); nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if reg.UUID == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		err := StoreRequest(reg)
		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	default:
		NotImplemented(w)
	}
}
