package main

import (
	"golang-hotel-management/middleware"
	"golang-hotel-management/database"
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
	routes.userRoutes(router)
	router.Use(middleware.Authentication())

	routes.foodRoutes(router)
	routes.menuRoutes(router)
	routes.tableRoutes(router)
	routes.orderRoutes(router)
	routes.orderItemsRoutes(router)
	routes.invoiceRoutes(router)

	router.Run(":" + port)

	_, err := .Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}
