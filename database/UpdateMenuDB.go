package database

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

func UpdateMenuDB(c *gin.Context, incomingMenu models.UpdateMenu) (models.UpdateMenu, error) {

	query := `UPDATE menus
		SET name = $2, category = $3, updated_at = Now()
		WHERE id = $1
		RETURNING id, name, category, updated_at`

	var updatedMenu models.UpdateMenu
	err := DbConn.QueryRow(query, incomingMenu.ID, incomingMenu.Name, incomingMenu.Category).Scan(
		&updatedMenu.ID, &updatedMenu.Name, &updatedMenu.Category, &updatedMenu.UpdatedAt,
	)
	if err != nil {
		return models.UpdateMenu{}, err
	}

	return updatedMenu, nil
}
