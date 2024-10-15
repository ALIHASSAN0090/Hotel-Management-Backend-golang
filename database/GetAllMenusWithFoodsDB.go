package database

import (
	"encoding/json"
	"net/http"

	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

func GetAllMenusWithFoodsDB(c *gin.Context) (*[]models.MenuItems, error) {
	var MenuItems []models.MenuItems

	query := `SELECT menus.id, menus.name, 
                     json_agg(json_build_object('id', food_items.id, 'name', food_items.name)) AS food_items
              FROM menus
              JOIN food_items ON menus.id = food_items.menu_id
              GROUP BY menus.id, menus.name`

	rows, err := DbConn.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var MenuItem models.MenuItems
		var foodItems string

		if err := rows.Scan(&MenuItem.ID, &MenuItem.Name, &foodItems); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return nil, err
		}

		if err := json.Unmarshal([]byte(foodItems), &MenuItem.Foods); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return nil, err
		}

		MenuItems = append(MenuItems, MenuItem)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil, err
	}

	return &MenuItems, nil
}
