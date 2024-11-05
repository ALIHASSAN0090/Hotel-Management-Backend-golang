package database

import (
	"encoding/json"
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrdersDB(c *gin.Context) ([]models.AllOrders, error) {

	query := `SELECT 
    o.id AS Order_Id, 
    o.total_price, 
    o.status,  -- Include the status column
    json_agg(
        json_build_object(
            'food_item_id', ofi.food_item_id,
            'name', f.name,
            'price_per_unit', f.price,
            'food_category', m.name,
            'category', m.category
        )
    ) AS food_items
FROM 
    orders AS o
JOIN 
    order_food_items AS ofi ON o.id = ofi.order_id
JOIN 
    food_items AS f ON ofi.food_item_id = f.id
JOIN 
    menus AS m ON f.id = m.id
GROUP BY 
    o.id, o.total_price, o.status; `

	rows, err := DbConn.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil, err
	}
	defer rows.Close()

	var orders []models.AllOrders

	for rows.Next() {
		var order models.AllOrders
		var foodItems string

		if err := rows.Scan(&order.OrderID, &order.TotalPrice, &order.Status, &foodItems); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return nil, err
		}

		if err := json.Unmarshal([]byte(foodItems), &order.FoodItems); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil, err
	}

	return orders, nil
}
