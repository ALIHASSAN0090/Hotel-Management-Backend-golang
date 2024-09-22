package database

import (
	"errors"
	"golang-hotel-management/models"
)

func GetInvoiceDB(order_id int64) (models.Invoice, error) {

	if order_id == 0 {
		return models.Invoice{}, errors.New("order_id is required")
	}

	query := `SELECT * FROM invoices WHERE order_id = $1`

	row := DbConn.QueryRow(query, order_id)

	var invoice models.Invoice
	err := row.Scan(&invoice.ID, &invoice.OrderID, &invoice.PaymentMethod, &invoice.PaymentStatus, &invoice.CreatedAt, &invoice.UpdatedAt)
	if err != nil {
		return models.Invoice{}, err

	}

	return invoice, nil
}
