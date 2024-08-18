package routes

import (
	"golang-hotel-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderitems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.POST("/orderItems/orderItem_id", controllers.UpdateOrderItem())
}
