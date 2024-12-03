package database

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

type invoiceRepository struct{}

func NewInvoiceRepository() *invoiceRepository {
	return &invoiceRepository{}
}

func (ir *invoiceRepository) CreateInvoiceDB(order_id int64) (models.CreateInvoice, error) {
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

func (ir *invoiceRepository) GetAllInvoicesDB(c *gin.Context) ([]models.Invoice, error) {

	query := `SELECT * FROM invoices`

	rows, err := DbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var invoice models.Invoice
		var createdAt sql.NullTime
		var updatedAt sql.NullTime

		if err := rows.Scan(&invoice.ID, &invoice.OrderID, &invoice.PaymentMethod, &invoice.PaymentStatus, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		if createdAt.Valid {
			invoice.CreatedAt = createdAt.Time
		}

		if updatedAt.Valid {
			invoice.UpdatedAt = updatedAt.Time
		}

		invoices = append(invoices, invoice)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (ir *invoiceRepository) GetInvoiceDB(order_id int64) (models.Invoice, error) {

	if order_id == 0 {
		return models.Invoice{}, errors.New("order_id is required")
	}

	query := `SELECT * FROM invoices WHERE order_id = $1`

	row := DbConn.QueryRow(query, order_id)

	var invoice models.Invoice
	var createdAt sql.NullTime
	var updatedAt sql.NullTime

	err := row.Scan(&invoice.ID, &invoice.OrderID, &invoice.PaymentMethod, &invoice.PaymentStatus, &createdAt, &updatedAt)
	if err != nil {
		return models.Invoice{}, err
	}

	if createdAt.Valid {
		invoice.CreatedAt = createdAt.Time
	}

	if updatedAt.Valid {
		invoice.UpdatedAt = updatedAt.Time
	}

	return invoice, nil
}

func (ir *invoiceRepository) UpdateInvoice(incomingInvoice models.UpdateInvoice) (models.Invoice, error) {

	query := `UPDATE invoices
	SET payment_method = $1, payment_status = $2, updated_at = NOW()
	WHERE order_id = $3
	RETURNING id, order_id, payment_method, payment_status, created_at, updated_at`

	row := DbConn.QueryRow(query, incomingInvoice.PaymentMethod, incomingInvoice.PaymentStatus, incomingInvoice.ID)

	var updatedInvoice models.Invoice
	err := row.Scan(&updatedInvoice.ID, &updatedInvoice.OrderID, &updatedInvoice.PaymentMethod, &updatedInvoice.PaymentStatus, &updatedInvoice.CreatedAt, &updatedInvoice.UpdatedAt)
	if err != nil {
		return models.Invoice{}, err
	}

	return updatedInvoice, nil
}
