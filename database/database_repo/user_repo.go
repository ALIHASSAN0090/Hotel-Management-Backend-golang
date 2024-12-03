package database_repo

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(c *gin.Context, User models.User) (models.User, error)
	GetUserByEmailDB(c *gin.Context, email string) (models.User, error)
	GetUserDB(c *gin.Context, userId int64) (models.AllUsers, error)
	GetUsersDB(c *gin.Context) ([]models.AllUsers, error)
	GetHotelDetailsDB(c *gin.Context) (string, error)
}
