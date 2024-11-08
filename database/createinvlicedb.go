package database

import (
	"fmt"
	"golang-hotel-management/models"
)

func CreateInvoiceDB(order_id int64) (models.CreateInvoice, error) {
	fmt.Println("Creating invoice for order_id:", order_id)

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM orders WHERE id = $1)`
	err := DbConn.QueryRow(checkQuery, order_id).Scan(&exists)
	if err != nil {
		return models.CreateInvoice{}, err
	}
	if !exists {
		return models.CreateInvoice{}, fmt.Errorf("order_id %d does not exist in orders table", order_id)
	}

	query := `INSERT INTO invoices(order_id, payment_status, created_at)
	VALUES ($1, $2, CURRENT_TIMESTAMP) RETURNING order_id, payment_status`
	var data models.CreateInvoice
	err = DbConn.QueryRow(query, order_id, data.PaymentStatus).Scan(
		&data.OrderID, &data.PaymentStatus,
	)
	if err != nil {
		return models.CreateInvoice{}, err
	}
	return data, nil
}
