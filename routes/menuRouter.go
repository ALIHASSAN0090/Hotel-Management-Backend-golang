package routes

import (
	"golang-hotel-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.GetAllMenusWithFoods())
	incomingRoutes.GET("/menus/:menu_id", controllers.GetAllFoods())
	incomingRoutes.POST("/menus/create", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/update", controllers.UpdateMenu())
}
