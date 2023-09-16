package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/twentymls/go-server-test/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	user, error := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: params.Name,
	})

	if error != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to create user: %v", error))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	user, error := apiCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:   user.ID,
		Name: params.Name,
	})

	if error != nil {
		respondWithError(w, 422, fmt.Sprintf("Unprocessable entity: %v", error))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {

	error := apiCfg.DB.DeleteUser(r.Context(), user.ID)

	if error != nil {
		respondWithError(w, 422, fmt.Sprintf("Unprocessable entity: %v", error))
		return
	}

	respondWithJSON(w, 200, fmt.Sprintf("User %v deleted", user.ID))
}

func (apiCfg *apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {

	users, error := apiCfg.DB.GetUsers(r.Context())

	if error != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to load users: %v", error))
		return
	}

	respondWithJSON(w, 200, databaseUsersToUsers(users))
}
