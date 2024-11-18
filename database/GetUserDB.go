package database

import (
	"database/sql"
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserDB(c *gin.Context, userId int64) (models.AllUsers, error) {
	query := `SELECT 
    u.id, 
    u.username, 
    u.email, 
    u.created_at, 
    COALESCE(SUM(CASE WHEN LOWER(i.payment_status) IN ('paid', 'Paid', 'completed') THEN o.total_price ELSE 0 END), 0) AS total_price_spent_on_orders_and_reservations,
    COUNT(DISTINCT CASE WHEN LOWER(i.payment_status) IN ('paid','Paid' ,'completed') THEN o.id ELSE NULL END) AS number_of_orders
FROM 
    users AS u
LEFT JOIN 
    orders AS o ON u.id = o.user_id
LEFT JOIN 
    invoices AS i ON o.id = i.order_id
WHERE 
    u.id = $1
GROUP BY 
    u.id, u.username, u.email, u.created_at;`

	row := DbConn.QueryRow(query, userId)

	var user models.AllUsers
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.Total_price_spent_on_orders_and_reservations, &user.NumberOfOrders); err != nil {
		if err == sql.ErrNoRows {
			return models.AllUsers{}, nil
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return models.AllUsers{}, err
	}

	return user, nil
}
