package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, error := json.Marshal(payload)

	if error != nil {
		log.Println("failed to marshal response: %v", payload)
		w.WriteHeader(code)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
