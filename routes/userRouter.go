package routes

import (
	"golang-hotel-management/controllers"
	"golang-hotel-management/database"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userRepo := database.NewUserRepository()
	userController := controllers.NewUserController(userRepo)
	incomingRoutes.GET("/users", userController.GetUsers())
	incomingRoutes.GET("/user/:user_id", userController.GetUser())
	incomingRoutes.POST("/users/Login", userController.Login())
	incomingRoutes.POST("/users/Signup", userController.Signup())
	incomingRoutes.GET("/header/check", userController.CheckHeader())
}
