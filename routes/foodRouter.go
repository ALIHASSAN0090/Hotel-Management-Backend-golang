package routes

import (
	"golang-hotel-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetAllFoods())
	incomingRoutes.GET("/foods/food_id", controllers.GetFood())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PATCH("/food/update", controllers.UpdateFood())
}
