package main

import (
	"log"
	"squawker-backend/internal/api"
	"squawker-backend/internal/database"
)

func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	router := api.SetupRouter()

	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
