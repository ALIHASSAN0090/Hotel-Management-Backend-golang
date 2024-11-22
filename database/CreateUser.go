package database

import (
	"golang-hotel-management/models"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, User models.User) (models.User, error) {

	query := `INSERT INTO users (username, email, password_hash, first_name, last_name) 
VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at;`
	err := DbConn.QueryRow(query, User.Username, User.Email, User.PasswordHash, User.FirstName, User.LastName).Scan(&User.ID, &User.CreatedAt)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return models.User{}, err
	}
	return User, nil
}
