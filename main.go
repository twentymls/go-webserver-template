package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not found!")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Println("Server running on port", port)
	error := server.ListenAndServe()

	if error != nil {
		log.Fatal(error)
	}
}
