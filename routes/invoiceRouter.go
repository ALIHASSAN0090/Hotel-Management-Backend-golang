package routes

import (
	"golang-hotel-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices/all", controllers.GetAllInvoices())
	incomingRoutes.GET("/invoices/:order_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoices/create", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
