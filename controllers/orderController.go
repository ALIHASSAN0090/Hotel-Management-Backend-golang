package controllers

import (
	"golang-hotel-management/database"
	"golang-hotel-management/models"
	"golang-hotel-management/pdf"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := database.GetAllOrdersDB(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "All Orders Fetched Succesfully",
			Status:  200,
			Data:    data,
		})
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}
		data, err := database.GetOrderDB(c, order_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Order Fetched Successfully",
			Status:  200,
			Data:    data,
		})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		var createOrder models.CombinedOrderReservation
		if err := c.ShouldBindJSON(&createOrder); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		OrderId, reservationId, err := database.CreateOrderDB(c, createOrder)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		InvoiceData, err := database.CreateInvoiceDB(OrderId)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		Data, err := database.GetOrderDB(c, OrderId)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		if reservationId != 0 {
			Reservation, err := database.GetReservationDB(c, reservationId)
			if err != nil {
				c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
				return
			}

			err = pdf.GenerateAndSendPDF("pdf/reservation.html", "pdf/invoice2.pdf")

			c.JSON(http.StatusOK, models.Response{
				Message: "Order , invoice and Reservation Created and Fetched Successfully",
				Status:  200,
				Data: gin.H{
					"order":       Data,
					"reservation": Reservation,
					"Invoice":     InvoiceData,
				},
			})
		} else {
			c.JSON(http.StatusOK, models.Response{
				Message: "Order and Invoice Created and Fetched Successfully",
				Status:  200,
				Data: gin.H{
					"order": Data,

					"Invoice": InvoiceData,
				},
			})

		}
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("order_id"), 10, 64)

		data, err := database.UpdateOrderDB(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in Getting Data From Database"})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Order Created and fetched Succesfully",
			Status:  200,
			Data:    data,
		})

	}
}
