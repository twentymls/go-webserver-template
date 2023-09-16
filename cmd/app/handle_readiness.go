package main

import (
	"net/http"

	"github.com/twentymls/go-server-test/cmd/internal/http_response"
)

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	http_response.RespondWithJSON(w, 200, struct{}{})
}
