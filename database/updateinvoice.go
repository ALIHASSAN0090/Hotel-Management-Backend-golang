package database

import (
	"golang-hotel-management/models"
	"time"
)

func UpdateInvoice(incomingInvoice models.UpdateInvoice) (models.Invoice, error) {

	query := `UPDATE invoices
	SET  payment_method = $1, payment_status = $2, updated_at = $3
	WHERE id = $4
	RETURNING id, order_id, payment_method, payment_status, created_at, updated_at`

	row := DbConn.QueryRow(query, incomingInvoice.PaymentMethod, incomingInvoice.PaymentStatus, time.Now().UTC(), incomingInvoice.ID)

	var updatedInvoice models.Invoice
	err := row.Scan(&updatedInvoice.ID, &updatedInvoice.OrderID, &updatedInvoice.PaymentMethod, &updatedInvoice.PaymentStatus, &updatedInvoice.CreatedAt, &updatedInvoice.UpdatedAt)
	if err != nil {
		return models.Invoice{}, err
	}

	return updatedInvoice, nil
}
