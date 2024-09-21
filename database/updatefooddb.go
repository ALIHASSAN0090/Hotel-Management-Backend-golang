package database

import (
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateFoodDB(c *gin.Context, incomingfooddata models.UpdateFood) error {

	query := `UPDATE food_items
	SET name = $2, price = $3, menu_id = $4, updated_at = CURRENT_TIMESTAMP
	WHERE id = $1`

	_, err := DbConn.Exec(query, incomingfooddata.ID, incomingfooddata.Name, incomingfooddata.Price, incomingfooddata.MenuID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return err
	}

	return nil
}
