package routes

import (
	"golang-hotel-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetAllInvoices())
	incomingRoutes.GET("/invoices/invoice_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoices/create", controllers.CreateInvoice())
	incomingRoutes.POST("/invoices/invoice_id", controllers.UpdateInvoice())
}
