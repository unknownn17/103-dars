package main

import (
	"fmt"
	"log"
	"net/http"

	"fitness/internal/config"
	"fitness/internal/handler"
	"fitness/internal/storage"
	"fitness/pkg/logger"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Configuration failed to load: %v", err)
	}

	store := storage.NewMemoryStorage()
	userHandler := handler.NewUserHandler(store)

	http.HandleFunc("/users", handler.LogMiddleware(userHandler.UsersHandler))
	http.HandleFunc("/users/", handler.LogMiddleware(userHandler.UserHandler))

	port := cfg.GetString("server.port")
	logger.Info(fmt.Sprintf("Server is starting on port %s...", port))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}