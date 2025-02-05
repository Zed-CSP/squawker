package main

import (
	"log"
	"os"
	"squawker-backend/internal/api"
	"squawker-backend/internal/config"
	"squawker-backend/internal/database"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()
		
	// Get port from Heroku environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Setup and run the server
	router := api.SetupRouter(db)
	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
