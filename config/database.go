package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

// ConnectDatabase connects to PostgreSQL
func ConnectDatabase() {
	dbURL := os.Getenv("DATABASE_URL")
	var err error

	// Establish the connection
	DB, err = pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the database!")
}

// CloseDatabase closes the connection pool
func CloseDatabase() {
	DB.Close()
}
