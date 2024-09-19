package database

import (
	"fmt"
	"golang-hotel-management/models"
	"time"
)

func CreateFoodDB(incomingfoodmodel models.Food) error {

	// Logging input data for debugging
	fmt.Println("CreateFoodDB called with input: ", incomingfoodmodel)

	// Defining the SQL query
	query := `INSERT INTO food_items(name, price, menu_id, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5)`

	// Printing query to trace any errors in SQL syntax
	fmt.Println("SQL Query: ", query)

	// Adding current time for created_at and updated_at fields
	incomingfoodmodel.CreatedAt = time.Now()
	incomingfoodmodel.UpdatedAt = time.Now()

	// Executing the query and logging values
	fmt.Println("Executing query with values: ", incomingfoodmodel.Name, incomingfoodmodel.Price, incomingfoodmodel.MenuID, incomingfoodmodel.CreatedAt, incomingfoodmodel.UpdatedAt)

	// Executing the query
	result, err := DbConn.Exec(query,
		incomingfoodmodel.Name,
		incomingfoodmodel.Price,
		incomingfoodmodel.MenuID,
		incomingfoodmodel.CreatedAt,
		incomingfoodmodel.UpdatedAt,
	)

	// If there's an error, log the error
	if err != nil {
		fmt.Println("Error executing query: ", err)
		return err
	}

	// Printing query result for confirmation
	fmt.Println("Query executed successfully, result: ", result)

	return nil
}
