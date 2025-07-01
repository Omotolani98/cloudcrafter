package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

       _ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
       host := os.Getenv("POSTGRES_HOST")
       port := os.Getenv("POSTGRES_PORT")
       user := os.Getenv("POSTGRES_USER")
       password := os.Getenv("POSTGRES_PASSWORD")
       dbname := os.Getenv("POSTGRES_DB")
       connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
               host, port, user, password, dbname)
       var err error
       DB, err = sql.Open("postgres", connStr)
       if err != nil {
               log.Fatalf("Failed to connect to the database: %v", err)
       }
       err = DB.Ping()
       if err != nil {
               log.Fatalf("Database is not reachable: %v", err)
       }

	log.Println("Database connection established successfully")
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Failed to close the database connection: %v", err)
		}
	}
}
