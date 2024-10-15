package database

import "golang-hotel-management/models"

func UpdateInvoice(incomingInvoice models.Invoice, id int64) (models.Invoice, error) {

	query := `UPDATE invoices
	SET order_id = $1, payment_method = $2, payment_status = $3, updated_at = $4
	WHERE id = $5
	RETURNING id, order_id, payment_method, payment_status, created_at, updated_at`

	row := DbConn.QueryRow(query, incomingInvoice.OrderID, incomingInvoice.PaymentMethod, incomingInvoice.PaymentStatus, incomingInvoice.UpdatedAt, id)

	var updatedInvoice models.Invoice
	err := row.Scan(&updatedInvoice.ID, &updatedInvoice.OrderID, &updatedInvoice.PaymentMethod, &updatedInvoice.PaymentStatus, &updatedInvoice.CreatedAt, &updatedInvoice.UpdatedAt)
	if err != nil {
		return models.Invoice{}, err
	}

	return updatedInvoice, nil
}
