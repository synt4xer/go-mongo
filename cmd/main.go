package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	delivery "github.com/synt4xer/go-mongo/internal/delivery/http"
)

func main() {
	var handlers *delivery.UserHandler
	var err error

	err = godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	handlers, err = apiHandlers()

	if err != nil {
		log.Fatal(fmt.Errorf("could not wire application: %w", err))
	}

	startServer(handlers)
}

func startServer(handlers *delivery.UserHandler) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/api/v1/users", handlers.FindAll)
	router.Post("/api/v1/users", handlers.SaveUser)
	router.Get("/api/v1/users/{id}", handlers.FindById)
	router.Patch("/api/v1/users/{id}", handlers.UpdateUser)
	router.Delete("/api/v1/users/{id}", handlers.DeleteUser)

	port := ":" + os.Getenv("PORT")

	log.Printf("Server running at port %s", port)

	_ = http.ListenAndServe(port, router)
}
