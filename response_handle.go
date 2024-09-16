package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type responseErr struct {
	Err string `json:"error"`
}

func responseWithJSON(w http.ResponseWriter, payload interface{}, statusCode int) {
	res, err := json.Marshal(payload)
	if err != nil {
		statusCode = 500
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

func responseWithErr(w http.ResponseWriter, msg string, statusCode int) {
	var payload = responseErr{}
	payload.Err = msg
	res, err := json.Marshal(payload)
	if err != nil {
		statusCode = 500
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(res)
}
