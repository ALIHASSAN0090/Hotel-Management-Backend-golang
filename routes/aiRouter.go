package routes

import (
	"golang-hotel-management/controllers"

	"github.com/gin-gonic/gin"
)

func AIRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/chat/query", controllers.AiQuery())
}
