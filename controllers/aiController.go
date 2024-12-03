package controllers

import (
	"golang-hotel-management/database/database_repo"
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AiController struct {
	Repo database_repo.AiRepository
	Or   database_repo.OrderRepository
	Ur   database_repo.UserRepository
}

func NewAiController(repo database_repo.AiRepository, or database_repo.OrderRepository, ur database_repo.UserRepository) AiController {
	return AiController{
		Repo: repo,
		Or:   or,
		Ur:   ur,
	}
}
func (ac *AiController) AiQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		question := c.Query("question")
		userid, exist := c.Get("userID")
		if !exist {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting user id"})
			return
		}

		order_details, err := ac.Or.GetOrderByUserIdDB(userid, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting order details"})
			return
		}

		hotel_details, err := ac.Ur.GetHotelDetailsDB(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting hotel details"})
			return
		}

		responce, err := ac.Repo.GetAiResponceDB(order_details, hotel_details, question)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting AI response"})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Got Answer",
			Status:  http.StatusOK,
			Data:    responce,
		})
	}
}
