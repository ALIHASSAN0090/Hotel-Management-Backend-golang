package database_repo

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

type MenuRepository interface {
	GetAllMenusWithFoodsDB(c *gin.Context) (*[]models.MenuItems, error)
	CreateMenuDB(c *gin.Context, incomingMenu models.CreateMenu) (models.Menu, error)
	UpdateMenuDB(c *gin.Context, incomingMenu models.UpdateMenu) (models.UpdateMenu, error)
}
