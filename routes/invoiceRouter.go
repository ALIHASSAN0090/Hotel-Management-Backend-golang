package routes

import (
	"golang-hotel-management/controllers"
	"golang-hotel-management/database"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	invoiceRepo := database.NewInvoiceRepository()
	invoiceController := controllers.NewInvoiceController(invoiceRepo)
	incomingRoutes.GET("/invoices/all", invoiceController.GetAllInvoices())
	incomingRoutes.GET("/invoices/:order_id", invoiceController.GetInvoice())
	incomingRoutes.PATCH("/invoice/update", invoiceController.UpdateInvoice())
}
