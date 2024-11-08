package database

import (
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMenuDB(c *gin.Context, incomingMenu models.CreateMenu) (models.Menu, error) {

	var createdMenu models.Menu

	duplicationQuery := `SELECT id FROM menus WHERE name = $1 AND category = $2`
	var existingID int
	if err := DbConn.QueryRow(duplicationQuery, incomingMenu.Name, incomingMenu.Category).Scan(&existingID); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Menu item already exists"})
		return models.Menu{}, err
	}

	query := `INSERT INTO menus (name, category, created_at) VALUES ($1, $2, NOW()) RETURNING id, name, category, created_at`

	if err := DbConn.QueryRow(query, incomingMenu.Name, incomingMenu.Category).Scan(&createdMenu.ID, &createdMenu.Name, &createdMenu.Category, &createdMenu.CreatedAt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error in executing query"})
		return models.Menu{}, err
	}

	return createdMenu, nil
}
