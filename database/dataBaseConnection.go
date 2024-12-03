package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DbConn *sql.DB

func InitializeRepo() (DbConn *sql.DB) {
	DbConn, _ = Connect()
	return
}

func Connect() (*sql.DB, error) {
	var err error

	// Load environment variables from the .env file
	err = godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
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
		if err != nil {
			log.Printf("Failed to open connection: %v. Retrying... (%d/3)", err, i+1)
			time.Sleep(5 * time.Second)
			continue
		}

		err = DbConn.Ping()
		if err == nil {
			log.Println("Successfully connected to the database")
			return DbConn, nil
		}

		log.Printf("Failed to ping database: %v. Retrying... (%d/3)", err, i+1)
		time.Sleep(5 * time.Second)
	}

	return DbConn, fmt.Errorf("failed to connect to the database after multiple attempts: %w", err)
}
