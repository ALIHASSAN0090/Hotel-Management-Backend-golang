package controller_repo

import "github.com/gin-gonic/gin"

type InvoiceController interface {
	GetAllInvoices() gin.HandlerFunc
	GetInvoice() gin.HandlerFunc
	UpdateInvoice() gin.HandlerFunc
}
