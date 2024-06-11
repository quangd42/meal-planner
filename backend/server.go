package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/quangd42/meal-planner/backend/handlers"

	_ "github.com/lib/pq"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	r := chi.NewRouter()
	handlers.AddRoutes(r)

	fmt.Printf("listening on port %s...\n", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
	return nil
}
