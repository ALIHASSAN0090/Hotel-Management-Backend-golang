package database

import (
	"encoding/json"
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderDB(c *gin.Context, incmoingOrderId int64) (*models.AllOrders, error) {

	query := `SELECT 
    o.id AS Order_Id, 
    o.total_price, 
    o.status, 
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
WHERE 
    o.id = $1
GROUP BY 
    o.id, o.total_price, o.status;`

	row := DbConn.QueryRow(query, incmoingOrderId)

	var order models.AllOrders
	var foodItems string

	if err := row.Scan(&order.OrderID, &order.TotalPrice, &order.Status, &foodItems); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil, err
	}

	if err := json.Unmarshal([]byte(foodItems), &order.FoodItems); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil, err
	}

	return &order, nil
}
