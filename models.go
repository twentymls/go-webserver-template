package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/twentymls/go-server-test/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(databaseUser database.User) User {
	return User{
		ID:        databaseUser.ID,
		CreatedAt: databaseUser.CreatedAt,
		UpdatedAt: databaseUser.UpdatedAt,
		Name:      databaseUser.Name,
		ApiKey:    databaseUser.ApiKey,
	}
}
