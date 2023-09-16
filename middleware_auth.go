package main

import (
	"fmt"
	"net/http"

	"github.com/twentymls/go-server-test/internal/auth"
	"github.com/twentymls/go-server-test/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("started GET v1/users...")

		apiKey, error := auth.GetApiKey(r.Header)

		if error != nil {
			respondWithError(w, 403, fmt.Sprintf("Failed to get API key: %v", error))
			return
		}

		user, error := cfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if error != nil {
			respondWithError(w, 401, fmt.Sprintf("Unauthorized"))
			return
		}

		handler(w, r, user)
	}
}
