package database

import (
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMenuDB(c *gin.Context, incomingMenu models.CreateMenu) error {

	query := `INSERT INTO menus (name , category , created_at) values ($1,$2,NOW())`

	if _, err := DbConn.Exec(query, incomingMenu.Name, incomingMenu.Category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error in executing query"})
		return err
	}
	return nil
}
