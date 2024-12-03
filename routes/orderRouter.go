package routes

import (
	"golang-hotel-management/controllers"
	"golang-hotel-management/database"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	orderRepo := database.NewOrderRepository()
	invoiceRepo := database.NewInvoiceRepository()
	orderController := controllers.NewOrderController(orderRepo, invoiceRepo)
	incomingRoutes.GET("/orders", orderController.GetOrders())
	incomingRoutes.GET("/orders/:order_id", orderController.GetOrder())
	incomingRoutes.POST("/orders", orderController.CreateOrder())
	incomingRoutes.POST("/update/order/:order_id", orderController.UpdateOrder())
}
