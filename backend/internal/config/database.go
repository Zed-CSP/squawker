package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetDatabaseURL() string {
	// Heroku provides DATABASE_URL, fallback to local config
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Local development database URL
		dbURL = "postgres://postgres:postgres@localhost:5432/squawker?sslmode=disable"
	}
	return dbURL
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", GetDatabaseURL())
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
