package database

import (
	"fmt"
	"golang-hotel-management/models"
)

func CreateFoodDB(incomingfoodmodel models.CreateFood) error {

	query := `INSERT INTO food_items(name, price, menu_id, created_at)
			  VALUES ($1, $2, $3, $4)`

	fmt.Println("Executing query with values: ", incomingfoodmodel.Name, incomingfoodmodel.Price, incomingfoodmodel.MenuID, incomingfoodmodel.CreatedAt)

	_, err := DbConn.Exec(query,
		incomingfoodmodel.Name,
		incomingfoodmodel.Price,
		incomingfoodmodel.MenuID,
		incomingfoodmodel.CreatedAt,
	)

	if err != nil {
		fmt.Println("Error executing query: ", err)
		return err
	}

	return nil
}
