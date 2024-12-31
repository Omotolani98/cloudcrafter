package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is a global database connection pool
var DB *sql.DB

// Init initializes the database connection
func Init() {
	// Read environment variables
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Build connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to the database
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Ping the database to ensure the connection is established
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	log.Println("Database connection established successfully")
}

// Close closes the database connection
func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Failed to close the database connection: %v", err)
		}
	}
}
