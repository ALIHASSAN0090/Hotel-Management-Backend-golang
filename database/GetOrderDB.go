package database

import (
	"encoding/json"
	"fmt"
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderDB(c *gin.Context, incmoingOrderId int64) (*models.AllOrders, error) {

	query := `SELECT 
    o.id AS Order_Id, 
    o.total_price, 
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
    o.id, o.total_price;`

	row := DbConn.QueryRow(query, incmoingOrderId)

	var order models.AllOrders
	var foodItems string

	if err := row.Scan(&order.OrderID, &order.TotalPrice, &foodItems); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil, err
	}

	var paymentStatus string
	paymentQuery := `SELECT payment_status FROM invoices WHERE order_id = $1`
	if err := DbConn.QueryRow(paymentQuery, incmoingOrderId).Scan(&paymentStatus); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil, err
	}
	order.Status = paymentStatus

	if err := json.Unmarshal([]byte(foodItems), &order.FoodItems); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil, err
	}

	return &order, nil
}

func GetOrderByUserIdDB(userid int, c *gin.Context) (string, error) {

	query := `
    SELECT 
        o.id AS order_id, 
        o.total_price, 
        i.payment_status, 
        i.payment_method,
        r.dine_in_time,
        r.number_of_persons
    FROM 
        orders AS o
    JOIN 
        invoices AS i ON o.id = i.order_id
    JOIN 
        reservations AS r ON r.order_id = o.id
    WHERE 
        o.user_id = ?;  
	`

	rows, err := DbConn.Query(query, userid)
	if err != nil {
		return "", fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var result string

	for rows.Next() {
		var orderID, numberOfPersons int
		var totalPrice float64
		var paymentStatus, paymentMethod, dineInTime string

		err := rows.Scan(&orderID, &totalPrice, &paymentStatus, &paymentMethod, &dineInTime, &numberOfPersons)
		if err != nil {
			return "", fmt.Errorf("error scanning row: %v", err)
		}

		result += fmt.Sprintf("Order ID: %d\nTotal Price: %.2f\nPayment Status: %s\nPayment Method: %s\nDine-in Time: %s\nNumber of Persons: %d\n\n",
			orderID, totalPrice, paymentStatus, paymentMethod, dineInTime, numberOfPersons)
	}

	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("error during row iteration: %v", err)
	}

	if result == "" {
		return "No orders found for the user.", nil
	}

	return result, nil
}
