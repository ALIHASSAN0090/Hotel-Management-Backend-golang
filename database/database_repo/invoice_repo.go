package database_repo

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

type InvoiceRepository interface {
	CreateInvoiceDB(order_id int64) (models.CreateInvoice, error)
	GetAllInvoicesDB(c *gin.Context) ([]models.Invoice, error)
	GetInvoiceDB(order_id int64) (models.Invoice, error)
	UpdateInvoice(incomingInvoice models.UpdateInvoice) (models.Invoice, error)
}
