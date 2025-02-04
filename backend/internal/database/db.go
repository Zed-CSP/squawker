package database

import (
	"database/sql"
	"log"
	"squawker-backend/internal/config"
)

var DB *sql.DB

// Initialize database connection
func InitDB() error {
	var err error
	DB, err = sql.Open("postgres", config.GetDatabaseURL())
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Successfully connected to database")
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return DB
} 