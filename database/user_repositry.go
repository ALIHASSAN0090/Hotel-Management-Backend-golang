package database

import (
	"database/sql"
	"fmt"
	"golang-hotel-management/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (ur *userRepository) CreateUser(c *gin.Context, User models.User) (models.User, error) {

	query := `INSERT INTO users (username, email, password_hash, first_name, last_name) 
VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at;`
	err := DbConn.QueryRow(query, User.Username, User.Email, User.PasswordHash, User.FirstName, User.LastName).Scan(&User.ID, &User.CreatedAt)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return models.User{}, err
	}
	return User, nil
}

func (ur *userRepository) GetUserByEmailDB(c *gin.Context, email string) (models.User, error) {
	var user models.User
	query := `SELECT id , username , email , password_hash , role FROM users WHERE email = $1`

	row := DbConn.QueryRowContext(c, query, email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *userRepository) GetUserDB(c *gin.Context, userId int64) (models.AllUsers, error) {
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

func (ur *userRepository) GetUsersDB(c *gin.Context) ([]models.AllUsers, error) {
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

func (ur *userRepository) GetHotelDetailsDB(c *gin.Context) (string, error) {

	query := `SELECT id , opening_time, closing_time, holiday, home_delivery, reservations FROM hotel_info`

	rows, err := DbConn.Query(query)
	if err != nil {
		return "", fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var id int
		var opening_time, closing_time, holiday, home_delivery, reservations string

		err := rows.Scan(&id, &opening_time, &closing_time, &holiday, &home_delivery, &reservations)
		if err != nil {
			return "", fmt.Errorf("error scanning row: %v", err)
		}

		result += fmt.Sprintf("ID: %d, Opening Time: %s, Closing Time: %s, Holiday: %s, Home Delivery: %s, Reservations: %s\n",
			id, opening_time, closing_time, holiday, home_delivery, reservations)
	}

	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("error during row iteration: %v", err)
	}
	return result, nil
}
