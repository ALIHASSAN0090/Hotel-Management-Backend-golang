package controllers

import (
	"golang-hotel-management/database"
	"golang-hotel-management/models"
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
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Order Fetched Succesfully",
			Status:  200,
			Data:    data,
		})
	}
}
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
