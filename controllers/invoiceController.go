package controllers

import (
	controller_repo "golang-hotel-management/controllers/controllers_repo"
	"golang-hotel-management/database/database_repo"
	"golang-hotel-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	Repo database_repo.InvoiceRepository
}

func NewInvoiceController(repo database_repo.InvoiceRepository) controller_repo.InvoiceController {
	return &InvoiceController{Repo: repo}
}

func (ic *InvoiceController) GetAllInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := ic.Repo.GetAllInvoicesDB(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

func (ic *InvoiceController) GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id format"})
			return
		}

		data, err := ic.Repo.GetInvoiceDB(order_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

func (ic *InvoiceController) UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

		var invoice models.UpdateInvoice

		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		data, err := ic.Repo.UpdateInvoice(invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
