package routes

import (
	"golang-hotel-management/controllers"
	"golang-hotel-management/database"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	orderController := controllers.NewOrderController(database.NewOrderRepository(), database.NewInvoiceRepository())

	incomingRoutes.GET("/orders", orderController.GetOrders())
	incomingRoutes.GET("/orders/:order_id", orderController.GetOrder())
	incomingRoutes.POST("/orders", orderController.CreateOrder())
	incomingRoutes.POST("/update/order/:order_id", orderController.UpdateOrder())
}
