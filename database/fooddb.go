package database

import (
	"database/sql"

	"golang-hotel-management/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFoodByFoodIdDB(incomingFoodId int64, c *gin.Context) (data models.Food, err error) {

	query := `SELECT name, price , menu_id FROM food_items WHERE id = $1`
	var food models.Food

	err = DbConn.QueryRow(query, incomingFoodId).Scan(
		&food.Name,
		&food.Price,
		&food.MenuID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("Food with ID not found")
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Food not found"})
		} else {
			log.Printf("Error fetching food details: %v", err)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "An error occurred while fetching food details"})
		}
		return models.Food{}, err
	}

	return food, nil
}
