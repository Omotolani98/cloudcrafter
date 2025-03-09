package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Omotolani98/cloudcrafter/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5060"
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		log.Println("Shutting down server...")
		db.Close()
		os.Exit(0)
	}()

	log.Println("Starting server on :" + port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
