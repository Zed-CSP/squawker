package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DATABASE_URL")


	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	log.Printf("Attempting to connect to database...")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Successfully connected to database")
	return DB
}
