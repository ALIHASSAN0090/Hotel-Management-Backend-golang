package main

import (
	"golang-hotel-management/database"
	"golang-hotel-management/middleware"
	"golang-hotel-management/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		log.Println("Port not found in environment, using default port 3000")
		port = "3000"
	}
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.New(middleware.GetCORSConfig()))
	routes.UserRoutes(router)
	router.Use(middleware.AuthMiddleware())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.AIRoutes(router)

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	router.Run(":" + port)
}
