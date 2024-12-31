package main

import (
	"cloudcrafter/internal/db"
	"cloudcrafter/pkg/handlers"
	"cloudcrafter/pkg/providers"
	"cloudcrafter/pkg/services"
	"cloudcrafter/routes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Get the port from the environment variable, with a fallback to ":8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = "5060" // Default port
	}

	// Initialize the provider registry
	providerRegistry := providers.NewProviderRegistry()
	awsProvider, err := providers.NewAWSProvider("us-east-1")
	if err != nil {
		log.Fatalf("Failed to initialize AWS provider: %v", err)
	}
	providerRegistry.Register("aws", awsProvider)

	// Initialize the provisioning service and handler
	provisioningService := services.NewProvisioningService(providerRegistry)
	provisioningHandler := handlers.NewProvisioningHandler(provisioningService)

	// Register all routes
	routes.RegisterRoutes(router, provisioningHandler)

	// handle graceful shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		log.Println("Shutting down server...")
		db.Close()
		os.Exit(0)
	}()

	// Start the server
	log.Println("Starting server on :" + port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
