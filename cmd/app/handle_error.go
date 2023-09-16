package main

import (
	"net/http"

	"github.com/twentymls/go-server-test/cmd/internal/http_response"
)

func handleError(w http.ResponseWriter, r *http.Request) {
	http_response.RespondWithError(w, 404, "Something went wrong!")
}
