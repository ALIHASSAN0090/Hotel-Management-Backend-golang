package database

import (
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
		if err := rows.Scan(&invoice.ID, &invoice.OrderID, &invoice.PaymentMethod, &invoice.PaymentStatus, &invoice.CreatedAt, &invoice.UpdatedAt); err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}
