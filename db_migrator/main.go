package main

import "log"

func main() {
	if err := RunMigrations(); err != nil {
		log.Fatalf("Migrations Failed %v", err)
	}
	log.Println("Migrations applied successfully!")
}
