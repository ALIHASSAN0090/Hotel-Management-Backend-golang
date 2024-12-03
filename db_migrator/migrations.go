package main

import (
	"fmt"
	controller_repo "golang-hotel-management/controllers/controllers_repo"
	"golang-hotel-management/database"
	"golang-hotel-management/database/database_repo"
	"golang-hotel-management/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	UserController controller_repo.UserController
	UserRepository database_repo.UserRepository
}

func NewMigration(UserController controller_repo.UserController, UserRepository database_repo.UserRepository) Migration {
	return Migration{
		UserController: UserController,
		UserRepository: UserRepository,
	}
}
func (m *Migration) RunMigrations() error {
	db, err := database.Connect()
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create database driver: %w", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://db_migrator/migration_files",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Up migration failed: %v. Attempting to run down migrations.", err)
		if downErr := migration.Down(); downErr != nil {
			return fmt.Errorf("could not run down migrations after up failure: %w", downErr)
		}
		if upErr := migration.Up(); upErr != nil && upErr != migrate.ErrNoChange {
			return fmt.Errorf("could not run up migrations after down: %w", upErr)
		}
	}

	log.Println("Migrations aplied successfully!")

	if err := m.createDefaultAdminUser(); err != nil {
		return fmt.Errorf("failed to create default admin user: %w", err)
	}

	return nil
}
func (m *Migration) createDefaultAdminUser() error {
	// if m.UserController == nil {
	// 	return fmt.Errorf("UserController is not initialized")
	// }
	// if m.UserRepository == nil {
	// 	return fmt.Errorf("UserRepository is not initialized")
	// }

	adminEmail := os.Getenv("ADMIN_DEFAULT_EMAIL")
	adminPassword := os.Getenv("ADMIN_DEFAULT_PASSWORD")

	adminPasswordHash, err := m.UserController.HashPassword(adminPassword)
	if err != nil {
		return err
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

	_, err = m.UserRepository.CreateUser(c, adminUser)
	if err != nil {
		return fmt.Errorf("could not create admin user: %w", err)
	}

	return nil
}
