package main

import (
	"fmt"
	"golang-hotel-management/controllers"
	"golang-hotel-management/database"
	"golang-hotel-management/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {

	db, err := database.Connect()
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create database driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db_migrator/migration_files",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Up migration failed: %v. Attempting to run down migrations.", err)
		if downErr := m.Down(); downErr != nil {
			return fmt.Errorf("could not run down migrations after up failure: %w", downErr)
		}
		if upErr := m.Up(); upErr != nil && upErr != migrate.ErrNoChange {
			return fmt.Errorf("could not run up migrations after down: %w", upErr)
		}
	}

	log.Println("Migrations applied successfully!")

	if err := createDefaultAdminUser(); err != nil {
		return err
	}

	return nil
}

func createDefaultAdminUser() error {
	adminEmail := os.Getenv("ADMIN_DEFAULT_EMAIL")
	adminPassword := os.Getenv("ADMIN_DEFAULT_PASSWORD")

	adminPasswordHash, err := controllers.HashPassword(adminPassword)
	for {
		if err != nil {

			return err
		}

		if match, _ := controllers.VerifyPassword(adminPasswordHash, adminPassword); match {
			adminPassword = adminPasswordHash
			break
		}
	}
	if adminEmail == "" || adminPasswordHash == "" {
		return fmt.Errorf("environment variables ADMIN_EMAIL, ADMIN_PASSWORD, and ADMIN_ROLE must be set")
	}

	adminUser := models.User{
		Username:     "admin",
		Email:        adminEmail,
		PasswordHash: adminPasswordHash,
		FirstName:    "Admin",
		LastName:     "User",
		Role:         "admin",
	}

	c, _ := gin.CreateTestContext(nil)

	_, err = database.CreateUser(c, adminUser)
	if err != nil {
		return fmt.Errorf("could not create admin user: %w", err)
	}

	return nil
}
