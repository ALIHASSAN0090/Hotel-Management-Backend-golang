package controllers

import (
	"fmt"
	"golang-hotel-management/database"
	"golang-hotel-management/models"
	"golang-hotel-management/pdf"
	"log"
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

			// email data:
			customerName, existsName := c.Get("username")
			customerEmail, existsEmail := c.Get("email")
			if !existsName || !existsEmail {
				log.Printf("Missing customer information: Name: %v, Email: %v", customerName, customerEmail)
				c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Customer name or email not found in header"})
				return
			}

			customerNameStr, okName := customerName.(string)
			customerEmailStr, okEmail := customerEmail.(string)
			if !okName || !okEmail {
				log.Printf("Customer information is not of type string: Name: %v, Email: %v", customerName, customerEmail)
				c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Customer name or email is not of type string"})
				return
			}

			resDate := createOrder.MakeReservation.DineInDate.Format("2006-01-02")
			resTime := createOrder.MakeReservation.DineInTime.Format("15:04:05")
			numberOfPersons := createOrder.MakeReservation.NumberOfPersons
			foods := createOrder.CreateOrder.FoodItems_IDs

			totalPrice, _ := database.GetTotalPrice(foods)

			err = pdf.GenerateAndSendPDF("pdf/reservation.html", "pdf/invoice2.pdf", customerNameStr, customerEmailStr, resDate, resTime, numberOfPersons, totalPrice)
			if err != nil {
				fmt.Println("Error generating and sending PDF:", err)
			}

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
