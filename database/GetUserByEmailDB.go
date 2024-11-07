package database

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

func GetUserByEmailDB(c *gin.Context, email string) (models.User, error) {
	var user models.User
	query := `SELECT id , username , email , password_hash , role FROM users WHERE email = $1`

	row := DbConn.QueryRowContext(c, query, email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
