package controllers

import (
	"fmt"
	"golang-hotel-management/database/database_repo"
	"golang-hotel-management/models"
	pdf "golang-hotel-management/pdf/generate-pdf"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	Repo database_repo.OrderRepository
	IC   database_repo.InvoiceRepository
}

func NewOrderController(repo database_repo.OrderRepository, ic database_repo.InvoiceRepository) OrderController {
	return OrderController{
		Repo: repo,
		IC:   ic,
	}
}

func (oc *OrderController) GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := oc.Repo.GetAllOrdersDB(c)
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

func (oc *OrderController) GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}
		data, err := oc.Repo.GetOrderDB(c, order_id)
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

func (oc *OrderController) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		var createOrder models.CombinedOrderReservation
		if err := c.ShouldBindJSON(&createOrder); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		OrderId, reservationId, err := oc.Repo.CreateOrderDB(c, createOrder)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		InvoiceData, err := oc.IC.CreateInvoiceDB(OrderId)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		Data, err := oc.Repo.GetOrderDB(c, OrderId)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		if reservationId != 0 {
			Reservation, err := oc.Repo.GetReservationDB(c, reservationId)
			if err != nil {
				c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
				return
			}

			go func() {
				if err := oc.PrepareAndSendReservationEmail(c, createOrder); err != nil {
					c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
					return
				}
			}()

			c.JSON(http.StatusOK, models.Response{
				Message: "Order, invoice, and Reservation Created and Fetched Successfully",
				Status:  200,
				Data: gin.H{
					"order":       Data,
					"reservation": Reservation,
					"Invoice":     InvoiceData,
				},
			})
		} else {
			go func() {
				if err := oc.PrepareAndSendOrderEmail(c, createOrder); err != nil {
					c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
					return
				}
			}()
			c.JSON(http.StatusOK, models.Response{
				Message: "Order and Invoice Created and Fetched Successfully",
				Status:  200,
				Data: gin.H{
					"order":   Data,
					"Invoice": InvoiceData,
				},
			})
		}
	}
}

func (oc *OrderController) PrepareAndSendReservationEmail(c *gin.Context, createOrder models.CombinedOrderReservation) error {
	customerName, existsName := c.Get("username")
	customerEmail, existsEmail := c.Get("email")
	if !existsName || !existsEmail {
		log.Printf("Missing customer information: Name: %v, Email: %v", customerName, customerEmail)
		return fmt.Errorf("customer name or email not found in header")
	}

	customerNameStr, okName := customerName.(string)
	customerEmailStr, okEmail := customerEmail.(string)
	if !okName || !okEmail {
		log.Printf("Customer information is not of type string: Name: %v, Email: %v", customerName, customerEmail)
		return fmt.Errorf("customer name or email is not of type string")
	}

	resDate := createOrder.MakeReservation.DineInDate.Format("2006-01-02")
	resTime := createOrder.MakeReservation.DineInTime.Format("15:04:05")
	numberOfPersons := createOrder.MakeReservation.NumberOfPersons
	foods := createOrder.CreateOrder.FoodItems_IDs

	totalPrice, _ := oc.Repo.GetTotalPrice(foods)
	totalFoods, _ := oc.Repo.GetOrderFoodsDB(foods)
	err := pdf.GenerateAndSendReservationPDF("pdf/files/reservation.html", "pdf/files/invoice2.pdf", customerNameStr, customerEmailStr, resDate, resTime, numberOfPersons, totalPrice, totalFoods)
	if err != nil {
		fmt.Println("Error generating and sending PDF:", err)
		return err
	}

	return nil
}

func (oc *OrderController) UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("order_id"), 10, 64)

		data, err := oc.Repo.UpdateOrderDB(c, id)
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

func (oc *OrderController) PrepareAndSendOrderEmail(c *gin.Context, order models.CombinedOrderReservation) error {
	customerName, existsName := c.Get("username")
	customerEmail, existsEmail := c.Get("email")
	if !existsName || !existsEmail {
		log.Printf("Missing customer information: Name: %v, Email: %v", customerName, customerEmail)
		return fmt.Errorf("customer name or email not found in header")
	}

	customerNameStr, okName := customerName.(string)
	customerEmailStr, okEmail := customerEmail.(string)
	if !okName || !okEmail {
		log.Printf("Customer information is not of type string: Name: %v, Email: %v", customerName, customerEmail)
		return fmt.Errorf("customer name or email is not of type string")
	}

	orderDate := time.Now().Format("2006-01-02")
	orderTime := time.Now().Format("15:04:05")

	totalPrice, _ := oc.Repo.GetTotalPrice(order.CreateOrder.FoodItems_IDs)
	totalFoods, _ := oc.Repo.GetOrderFoodsDB(order.CreateOrder.FoodItems_IDs)
	err := pdf.GenerateAndSendOrderPDF("pdf/files/order.html", "pdf/files/invoice2.pdf", customerNameStr, customerEmailStr, orderDate, orderTime, 0, totalPrice, totalFoods)
	if err != nil {
		fmt.Println("Error generating and sending PDF:", err)
		return err
	}

	return nil
}
