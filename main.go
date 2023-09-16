package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/twentymls/go-server-test/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not found!")
	}

	router := chi.NewRouter()

	dataBaseUrl := os.Getenv("DB_URL")

	if dataBaseUrl == "" {
		log.Fatal("DB_URL not found!")
	}

	connection, db_error := sql.Open("postgres", dataBaseUrl)
	if db_error != nil {
		log.Fatal(db_error)
	}

	apiCfg := apiConfig{
		DB: database.New(connection),
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handleReadiness)
	v1Router.Get("/error", handleError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	router.Mount("/v1", v1Router)

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
