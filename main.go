package main

import (
	"golang-hotel-management/database"
	"golang-hotel-management/middleware"
	"golang-hotel-management/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(middleware.AuthMiddleware())
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)

	routes.InvoiceRoutes(router)

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	router.Run(":3000")
}
