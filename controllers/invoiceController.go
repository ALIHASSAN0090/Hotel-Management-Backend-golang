package controllers

import (
	"fmt"
	"golang-hotel-management/database"
	"golang-hotel-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := database.GetAllInvoicesDB(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id format"})
			return
		}
		fmt.Println(order_id)
		data, err := database.GetInvoiceDB(order_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

		// var invoice models.Invoice

		// err := c.ShouldBindJSON(&invoice)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		// }

		// invoiceres, err := database.CreateInvoiceDB(invoice)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }

		// c.JSON(http.StatusOK, invoiceres)

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

		var invoice models.UpdateInvoice

		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		data, err := database.UpdateInvoice(invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
