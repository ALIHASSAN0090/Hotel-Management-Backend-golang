package routes

import (
	"golang-hotel-management/controllers"
	"golang-hotel-management/database"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	menuRepo := database.NewMenuRepository()
	menuController := controllers.NewMenuController(menuRepo)
	incomingRoutes.GET("/menus", menuController.GetAllMenusWithFoods())
	incomingRoutes.POST("/menus/create", menuController.CreateMenu())
	incomingRoutes.PATCH("/menus/update", menuController.UpdateMenu())
}
