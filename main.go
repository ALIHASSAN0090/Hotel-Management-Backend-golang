package main

import (
	"golang-hotel-management/database"
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
	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close() // Ensure the database connection is closed when the application exits

	router.Run(":3000")
}
