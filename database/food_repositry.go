package database

import (
	"database/sql"
	"fmt"
	"golang-hotel-management/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type foodRepository struct{}

func NewFoodRepository() *foodRepository {
	return &foodRepository{}
}

func (r *foodRepository) GetAllFoodsDB(menuID int64, c *gin.Context) ([]models.FoodItem, error) {
	var foods []models.FoodItem
	query := `SELECT id, name, price, menu_id FROM food_items WHERE menu_id = $1`

	rows, err := DbConn.Query(query, menuID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var food models.FoodItem
		if err := rows.Scan(&food.ID, &food.Name, &food.Price, &food.MenuID); err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return foods, nil
}

func (r *foodRepository) GetFoodByFoodIdDB(foodID int64, c *gin.Context) (*models.FoodItem, error) {
	query := `SELECT id, name, price, menu_id FROM food_items WHERE id = $1`
	var food models.FoodItem

	err := DbConn.QueryRow(query, foodID).Scan(&food.ID, &food.Name, &food.Price, &food.MenuID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("Food with ID not found")
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Food not found"})
		} else {
			log.Printf("Error fetching food details: %v", err)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "An error occurred while fetching food details"})
		}
		return nil, err
	}

	return &food, nil
}

func (r *foodRepository) CreateFoodDB(food models.CreateFood) error {
	query := `INSERT INTO food_items(name, price, menu_id, created_at)
			  VALUES ($1, $2, $3, NOW())`

	fmt.Println("Executing query with values: ", food.Name, food.Price, food.MenuID)

	_, err := DbConn.Exec(query, food.Name, food.Price, food.MenuID)
	if err != nil {
		fmt.Println("Error executing query: ", err)
		return err
	}

	return nil
}

func (r *foodRepository) UpdateFoodDB(c *gin.Context, food models.UpdateFood) error {
	query := `UPDATE food_items
	SET name = $2, price = $3, menu_id = $4, updated_at = CURRENT_TIMESTAMP
	WHERE id = $1`

	_, err := DbConn.Exec(query, food.ID, food.Name, food.Price, food.MenuID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return err
	}

	return nil
}
