package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func GetDatabaseURL() string {
	// Heroku provides DATABASE_URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Local development database URL
		dbURL = "postgres://postgres:postgres@localhost:5432/squawker_db?sslmode=disable"
	}

	// Heroku SSL
	if os.Getenv("ENVIRONMENT") == "production" {
		dbURL += "&sslmode=require"
	}

	return dbURL
}
