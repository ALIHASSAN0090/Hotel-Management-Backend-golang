package main

import (
	controller_repo "golang-hotel-management/controllers/controllers_repo"
	"golang-hotel-management/database/database_repo"
	"log"
)

func main() {

	var userController controller_repo.UserController
	var userRepository database_repo.UserRepository

	migration := NewMigration(userController, userRepository)

	if err := migration.RunMigrations(); err != nil {
		log.Fatalf("Migrations Failed %v", err)
	}

	log.Println("Migrations applied successfully!")
}
