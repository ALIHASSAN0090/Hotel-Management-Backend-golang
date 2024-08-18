package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var DbConn *sql.DB

func Connect() (*sql.DB, error) {
	var err error

	// Load environment variables from .env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" || dbPassword == "" {
		return nil, fmt.Errorf("missing one or more required environment variables")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	for i := 0; i < 3; i++ {
		DbConn, err = sql.Open("postgres", psqlInfo)
		if err == nil {
			err = DbConn.Ping()
			if err == nil {
				break
			}
		}
		fmt.Printf("Failed to connect to the database. Retrying... (%d/3)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	} else {
		return nil, fmt.Errorf("Succesfully Connected to database: %v", err)
	}

	return DbConn, nil
}
