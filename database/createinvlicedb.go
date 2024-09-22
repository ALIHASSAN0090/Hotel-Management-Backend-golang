package database

import "golang-hotel-management/models"

func CreateInvoiceDB(data models.Invoice) (models.Invoice, error) {

	query := `INSERT INTO invoices(order_id, payment_method, payment_status, created_at)
	VALUES ($1, $2, $3, $4) RETURNING order_id, payment_method, payment_status, created_at`

	err := DbConn.QueryRow(query, data.OrderID, data.PaymentMethod, data.PaymentStatus, data.CreatedAt).Scan(
		&data.OrderID, &data.PaymentMethod, &data.PaymentStatus, &data.CreatedAt,
	)
	if err != nil {
		return models.Invoice{}, err
	}
	return data, nil
}
