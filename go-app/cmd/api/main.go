package main

import (
	"fmt"
	"log"

	"next-go-task/internal/config"
	"next-go-task/internal/repository"
	"next-go-task/internal/router"
	"next-go-task/internal/service"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize dependencies
	// TODO: Replace with actual database connection
	sampleRepo := repository.NewInMemorySampleRepository()
	sampleService := service.NewSampleService(sampleRepo)

	// Setup router (REST + GraphQL)
	e := router.Setup(cfg, sampleService)

	// Start server
	address := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on %s", address)
	log.Printf("GraphQL Playground: http://localhost:%s/playground", cfg.Port)
	log.Printf("REST API Health Check: http://localhost:%s/health", cfg.Port)

	if err := e.Start(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
