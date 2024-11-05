package database

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

func CreateOrderDB(c *gin.Context, createOrder models.CreateOrder) (int64, error) {

	totalPrice, err := GetTotalPrice(createOrder.FoodItems_IDs)
	if err != nil {
		return 0, err
	}

	var ID int64
	query1 := `INSERT INTO Orders (total_price, user_id, created_at) VALUES ($1, $2, CURRENT_DATE) RETURNING id`
	err = DbConn.QueryRow(query1, totalPrice, createOrder.UserID).Scan(&ID)
	if err != nil {
		return 0, err
	}

	for _, foodID := range createOrder.FoodItems_IDs {
		query2 := `INSERT INTO order_food_items (order_id, food_item_id) VALUES ($1, $2)`
		_, err := DbConn.Exec(query2, ID, foodID.ID)
		if err != nil {
			return 0, err
		}
	}

	return ID, nil
}

func GetTotalPrice(foodids []models.FoodID) (float64, error) {

	var totalPrice float64
	for _, foodid := range foodids {
		var price float64
		query := `SELECT price FROM food_items WHERE id = $1`
		err := DbConn.QueryRow(query, foodid.ID).Scan(&price)
		if err != nil {
			return 0.0, err
		}
		totalPrice += price
	}

	return totalPrice, nil
}
