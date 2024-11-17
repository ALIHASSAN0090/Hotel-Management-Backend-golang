package database

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

func GetUsersDB(c *gin.Context) ([]models.AllUsers, error) {
	var users []models.AllUsers

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
	GROUP BY 
		u.id, u.username, u.email, u.created_at;`

	rows, err := DbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.AllUsers
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.Total_price_spent_on_orders_and_reservations, &user.NumberOfOrders); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
