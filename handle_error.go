package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 404, "Something went wrong!")
}
