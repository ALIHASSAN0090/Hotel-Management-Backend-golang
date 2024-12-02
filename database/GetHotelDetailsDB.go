package database

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetHotelDetailsDB(c *gin.Context) (string, error) {

	query := `SELECT id , opening_time, closing_time, holiday, home_delivery, reservations FROM hotel_info`

	rows, err := DbConn.Query(query)
	if err != nil {
		return "", fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var id int
		var opening_time, closing_time, holiday, home_delivery, reservations string

		err := rows.Scan(&id, &opening_time, &closing_time, &holiday, &home_delivery, &reservations)
		if err != nil {
			return "", fmt.Errorf("error scanning row: %v", err)
		}

		result += fmt.Sprintf("ID: %d, Opening Time: %s, Closing Time: %s, Holiday: %s, Home Delivery: %s, Reservations: %s\n",
			id, opening_time, closing_time, holiday, home_delivery, reservations)
	}

	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("error during row iteration: %v", err)
	}
	return result, nil
}
