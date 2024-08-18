package routes

import (
	"golang-hotel-management/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTables())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.POST("/tables/table_id", controllers.UpdateTable())
}
