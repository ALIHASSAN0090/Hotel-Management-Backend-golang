package controllers

import (
	controller_repo "golang-hotel-management/controllers/controllers_repo"
	"golang-hotel-management/database/database_repo"
	"golang-hotel-management/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FoodController struct {
	Repo database_repo.FoodRepository
}

func NewFoodController(repo database_repo.FoodRepository) controller_repo.FoodController {
	return &FoodController{Repo: repo}
}

func (fc *FoodController) GetAllFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuID, err := strconv.ParseInt(c.Param("menu_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid menu ID"})
			return
		}

		allFoods, err := fc.Repo.GetAllFoodsDB(menuID, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error fetching foods from the database"})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Fetched Food Data",
			Status:  http.StatusOK,
			Data:    allFoods,
		})
	}
}

func (fc *FoodController) GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("food_id")
		if id == "" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Food ID is required"})
			return
		}

		foodID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid Food ID"})
			return
		}

		foodData, err := fc.Repo.GetFoodByFoodIdDB(int64(foodID), c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Error fetching food data"})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Fetched Food Data",
			Status:  http.StatusOK,
			Data:    foodData,
		})
	}
}

func (fc *FoodController) CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputFood models.CreateFood

		if err := c.ShouldBindJSON(&inputFood); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
			return
		}

		inputFood.CreatedAt = time.Now()

		if err := fc.Repo.CreateFoodDB(inputFood); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error creating food item"})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Food item created successfully",
			Status:  http.StatusOK,
			Data:    inputFood,
		})
	}
}

func (fc *FoodController) UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foodData models.UpdateFood

		if err := c.ShouldBindJSON(&foodData); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input"})
			return
		}

		if foodData.Name == "" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Name field is required"})
			return
		}
		if foodData.Price <= 0 {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Price must be a positive number"})
			return
		}
		if foodData.MenuID <= 0 {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "MenuID must be valid"})
			return
		}

		if err := fc.Repo.UpdateFoodDB(c, foodData); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error updating food item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Food item updated successfully",
			"data":    foodData,
		})
	}
}
