package routes

import (
	"golang-hotel-management/controllers"
	"golang-hotel-management/database"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {

	foodRepo := database.NewFoodRepository()
	foodController := controllers.NewFoodController(foodRepo)

	incomingRoutes.GET("/menus/:menu_id", foodController.GetAllFoods())
	incomingRoutes.GET("/foods/:food_id", foodController.GetFood())
	incomingRoutes.POST("/foods", foodController.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", foodController.UpdateFood())
}
