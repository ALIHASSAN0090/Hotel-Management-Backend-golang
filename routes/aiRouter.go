package routes

import (
	"golang-hotel-management/controllers"
	"golang-hotel-management/database"

	"github.com/gin-gonic/gin"
)

func AIRoutes(incomingRoutes *gin.Engine) {
	aiRepo := database.NewAiRepository()
	orderRepo := database.NewOrderRepository()
	userRepo := database.NewUserRepository()
	aiController := controllers.NewAiController(aiRepo, orderRepo, userRepo)
	incomingRoutes.POST("/chat/query", aiController.AiQuery())
}
