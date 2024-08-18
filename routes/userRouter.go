package routes

import (
	"golang-hotel-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUsers())
	incomingRoutes.POST("/users/Login", controllers.Login())
	incomingRoutes.POST("/users/Signup", controllers.Signup())
}
