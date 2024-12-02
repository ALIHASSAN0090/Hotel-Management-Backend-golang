package database

import (
	"database/sql"
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

func GetAllInvoicesDB(c *gin.Context) ([]models.Invoice, error) {

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
