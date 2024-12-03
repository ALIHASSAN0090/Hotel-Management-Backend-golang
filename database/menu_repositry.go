package database

import (
	"encoding/json"
	"net/http"

	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

type menuRepository struct{}

func NewMenuRepository() *menuRepository {
	return &menuRepository{}
}

func (mr *menuRepository) GetAllMenusWithFoodsDB(c *gin.Context) (*[]models.MenuItems, error) {
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

func (mr *menuRepository) CreateMenuDB(c *gin.Context, incomingMenu models.CreateMenu) (models.Menu, error) {

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

func (mr *menuRepository) UpdateMenuDB(c *gin.Context, incomingMenu models.UpdateMenu) (models.UpdateMenu, error) {

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
