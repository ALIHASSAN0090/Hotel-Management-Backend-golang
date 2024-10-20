package controllers

import (
	"golang-hotel-management/database"
	"golang-hotel-management/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		menu_id, err := strconv.ParseInt(c.Param("menu_id"), 10, 64)
		allfoods, err := database.GetAllFoodsDB(c, menu_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting all foods from database"})
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Fetched Food Data",
			Status:  http.StatusOK,
			Data:    allfoods,
		})
	}
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Query("food_id")
		log.Printf("Received request for food with ID: %s", id)

		if id == "" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Food ID is required"})
			return
		}

		idint, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error converting food_id to integer: %v", err)
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid Food ID"})
			return
		}
		var idint64 int64 = int64(idint)

		data, err := database.GetFoodByFoodIdDB(idint64, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error in getting data from db"})
			return
		}

		// Return the food data
		c.JSON(http.StatusOK, models.Response{
			Message: "Fetched Food Data",
			Status:  http.StatusOK,
			Data:    data,
		})
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputfood models.CreateFood

		if err := c.ShouldBindJSON(&inputfood); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid Input"})
			return
		}

		inputfood.CreatedAt = time.Now()

		if err := database.CreateFoodDB(inputfood); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error in creating food item in database"})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Food item created successfully",
			Status:  http.StatusOK,
			Data:    inputfood,
		})
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var fooddata models.UpdateFood

		if err := c.ShouldBindJSON(&fooddata); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input data"})
			return
		}

		if fooddata.Name == "" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Name field is required"})
			return
		}
		if fooddata.Price <= 0 {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Price must be a positive number"})
			return
		}
		if fooddata.MenuID <= 0 {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "MenuID must be valid"})
			return
		}

		if err := database.UpdateFoodDB(c, fooddata); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error updating food item in the database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Food item updated successfully",
			"data":    fooddata,
		})
	}
}

// func Round(num float64) int {

// }

// func Fixed(num float64, preceision int) float64 {

// }
