package main

import (
	"fmt"
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
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.tableRoutes(router)
	routes.orderRoutes(router)
	routes.orderItemsRoutes(router)
	routes.invoiceRoutes(router)

	router.Run(":" + port)

	_, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		fmt.Println("connected to the database")
	}
}
