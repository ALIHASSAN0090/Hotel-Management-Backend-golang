package database

import (
	"encoding/json"
	"fmt"
	"golang-hotel-management/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type orderRepository struct{}

func NewOrderRepository() *orderRepository {
	return &orderRepository{}
}

func (or *orderRepository) GetAllOrdersDB(c *gin.Context) ([]models.AllOrders, error) {

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

func (or *orderRepository) GetOrderDB(c *gin.Context, incmoingOrderId int64) (*models.AllOrders, error) {

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

func (or *orderRepository) GetOrderByUserIdDB(userid any, c *gin.Context) (string, error) {

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
        o.user_id = $1;  
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

func (or *orderRepository) UpdateOrderDB(c *gin.Context, updateId int64) (models.AllOrders, error) {
	query := `UPDATE orders
	SET status = 'paid'
	WHERE id = $1 RETURNING id, total_price, status`
	row := DbConn.QueryRow(query, updateId)
	var updatedOrder models.AllOrders
	err := row.Scan(&updatedOrder.OrderID, &updatedOrder.TotalPrice, &updatedOrder.Status)
	if err != nil {
		return models.AllOrders{}, err
	}

	return updatedOrder, nil
}

func (or *orderRepository) CreateOrderDB(c *gin.Context, createOrder models.CombinedOrderReservation) (int64, int64, error) {

	totalPrice, err := or.GetTotalPrice(createOrder.CreateOrder.FoodItems_IDs)
	if err != nil {
		return 0, 0, err
	}

	var ID int64
	query1 := `INSERT INTO Orders (total_price, user_id, created_at) VALUES ($1, $2, CURRENT_DATE) RETURNING id`
	err = DbConn.QueryRow(query1, totalPrice, createOrder.CreateOrder.UserID).Scan(&ID)
	if err != nil {
		return 0, 0, err
	}

	for _, foodID := range createOrder.CreateOrder.FoodItems_IDs {
		query2 := `INSERT INTO order_food_items (order_id, food_item_id) VALUES ($1, $2)`
		_, err := DbConn.Exec(query2, ID, foodID.ID)
		if err != nil {
			return 0, 0, err
		}
	}

	if createOrder.MakeReservation.NumberOfPersons == 0 || createOrder.MakeReservation.DineInDate.IsZero() || createOrder.MakeReservation.DineInTime.IsZero() {
		return ID, 0, nil
	} else {
		dineInDateTime := or.CalculateDineInDateTime(createOrder.MakeReservation.DineInDate, createOrder.MakeReservation.DineInTime)
		if dineInDateTime.IsZero() {
			return 0, 0, fmt.Errorf("invalid dine-in date and time")
		}

		query3 := `INSERT INTO reservations(order_id, number_of_persons, dine_in_time,  created_at) 
        VALUES ($1, $2, $3 , CURRENT_TIMESTAMP) RETURNING id`

		var resID int64
		err = DbConn.QueryRow(query3, ID, createOrder.MakeReservation.NumberOfPersons, dineInDateTime).Scan(&resID)
		if err != nil {
			return 0, 0, err
		}

		return ID, resID, nil

	}

}
func (or *orderRepository) GetTotalPrice(foodids []models.FoodID) (float64, error) {

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

func (or *orderRepository) GetOrderFoodsDB(foodids []models.FoodID) (string, error) {

	var foodNames []string
	for _, foodid := range foodids {
		var name string
		query := `SELECT name FROM food_items WHERE id = $1`
		err := DbConn.QueryRow(query, foodid.ID).Scan(&name)
		if err != nil {
			return "", err
		}
		foodNames = append(foodNames, name)
	}

	return strings.Join(foodNames, ", "), nil
}

func (or *orderRepository) CalculateDineInDateTime(dineInDate, dineInTime time.Time) time.Time {
	return dineInDate.Add(time.Hour*time.Duration(dineInTime.Hour()) + time.Minute*time.Duration(dineInTime.Minute()))
}

func (or *orderRepository) GetReservationDB(c *gin.Context, resId int64) (models.Reservation, error) {

	var reservation models.Reservation
	query := `SELECT order_id, number_of_persons, dine_in_time, created_at
	FROM reservations WHERE id = $1`
	err := DbConn.QueryRow(query, resId).Scan(&reservation.OrderID, &reservation.NumberOfPersons, &reservation.DineInTime, &reservation.CreatedAt)
	if err != nil {
		return models.Reservation{}, err
	}

	return reservation, nil
}
