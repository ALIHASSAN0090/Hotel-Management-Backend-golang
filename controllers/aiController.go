package controllers

import (
	"golang-hotel-management/database"
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AiQuery() gin.HandlerFunc {
	return func(c *gin.Context) {

		userid, exist := c.Get("userID")
		if !exist {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting question"})
		}
		question := c.Param("question")

		order_details, err := database.GetOrderByUserIdDB(userid, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting question"})
		}

		hotel_details, err := database.GetHotelDetailsDB(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting question"})
		}

		responce, err := database.GetAiResponceDB(order_details, hotel_details, question)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting question"})
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Got Answer",
			Status:  http.StatusOK,
			Data:    responce,
		})
	}
}
