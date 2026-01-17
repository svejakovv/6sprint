package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.Default()

	srv := server.New(logger)

	logger.Println("Starting server on :8080")
	if err := srv.Start(); err != nil {
		logger.Fatal("Server failed to start:", err)
	}
}
