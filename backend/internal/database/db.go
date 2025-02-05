package database

import (
	"database/sql"
	"log"
	"squawker-backend/internal/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error

	// Get database URL from config
	dbURL := config.GetDatabaseURL()

	// Open database connection
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to database")
	return db, nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}
