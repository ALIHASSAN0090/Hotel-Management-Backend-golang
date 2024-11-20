package main

import (
	"fmt"
	"golang-hotel-management/database"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	db, err := database.Connect()
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}
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
		return fmt.Errorf("could not run up migrations: %w", err)
	}

	log.Println("Migrations applied successfully!")
	return nil
}
