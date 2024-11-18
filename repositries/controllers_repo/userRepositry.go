package repo

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetUsersDB(c *gin.Context) ([]models.User, error)
	GetUserDB(c *gin.Context, userId int64) (models.User, error)
	GetUserByEmailDB(c *gin.Context, email string) (models.User, error)
	CreateUser(c *gin.Context, user models.User) (models.User, error)
}
