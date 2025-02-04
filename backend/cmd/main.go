package main

import (
	"log"
	"squawker-backend/internal/api"
)

func main() {
	router := api.SetupRouter()

	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
