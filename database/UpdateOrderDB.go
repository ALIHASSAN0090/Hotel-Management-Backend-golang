package database

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

func UpdateOrderDB(c *gin.Context, updateId int64) (models.AllOrders, error) {
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
