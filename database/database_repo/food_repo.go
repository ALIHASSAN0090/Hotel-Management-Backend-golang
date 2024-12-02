package database_repo

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

type FoodRepository interface {
	GetAllFoodsDB(menuID int64, c *gin.Context) ([]models.FoodItem, error)
	GetFoodByFoodIdDB(foodID int64, c *gin.Context) (*models.FoodItem, error)
	CreateFoodDB(food models.CreateFood) error
	UpdateFoodDB(c *gin.Context, food models.UpdateFood) error
}
