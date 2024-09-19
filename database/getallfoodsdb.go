package database

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

func GetAllFoodsDB(c *gin.Context) ([]models.Food, error) {
	var foods []models.Food
	query := `SELECT id, name, price, menu_id FROM food_items`

	rows, err := DbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var food models.Food
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
