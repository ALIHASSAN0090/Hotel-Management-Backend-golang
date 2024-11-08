package database

import (
	"fmt"
	"golang-hotel-management/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrderDB(c *gin.Context, createOrder models.CombinedOrderReservation) (int64, int64, error) {

	totalPrice, err := GetTotalPrice(createOrder.CreateOrder.FoodItems_IDs)
	if err != nil {
		return 0, 0, err
	}

	var ID int64
	query1 := `INSERT INTO Orders (total_price, user_id, created_at) VALUES ($1, $2, CURRENT_DATE) RETURNING id`
	err = DbConn.QueryRow(query1, totalPrice, createOrder.CreateOrder.UserID).Scan(&ID)
	if err != nil {
		return 0, 0, err
	}

	for _, foodID := range createOrder.CreateOrder.FoodItems_IDs {
		query2 := `INSERT INTO order_food_items (order_id, food_item_id) VALUES ($1, $2)`
		_, err := DbConn.Exec(query2, ID, foodID.ID)
		if err != nil {
			return 0, 0, err
		}
	}

	if createOrder.MakeReservation.NumberOfPersons == 0 || createOrder.MakeReservation.DineInDate.IsZero() || createOrder.MakeReservation.DineInTime.IsZero() {
		return ID, 0, nil
	} else {
		dineInDateTime := CalculateDineInDateTime(createOrder.MakeReservation.DineInDate, createOrder.MakeReservation.DineInTime)
		if dineInDateTime.IsZero() {
			return 0, 0, fmt.Errorf("invalid dine-in date and time")
		}

		formattedDineInTime := dineInDateTime.Format("2006-01-02 15:04:05")

		query3 := `INSERT INTO reservations(order_id, number_of_persons, dine_in_time, formatted_dine_in_time, created_at) 
        VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id`

		var resID int64
		err = DbConn.QueryRow(query3, ID, createOrder.MakeReservation.NumberOfPersons, dineInDateTime, formattedDineInTime).Scan(&resID)
		if err != nil {
			return 0, 0, err
		}

		

		return ID, resID, nil

	}

}
func GetTotalPrice(foodids []models.FoodID) (float64, error) {

	var totalPrice float64
	for _, foodid := range foodids {
		var price float64
		query := `SELECT price FROM food_items WHERE id = $1`
		err := DbConn.QueryRow(query, foodid.ID).Scan(&price)
		if err != nil {
			return 0.0, err
		}
		totalPrice += price
	}

	return totalPrice, nil
}

func CalculateDineInDateTime(dineInDate, dineInTime time.Time) time.Time {
	return dineInDate.Add(time.Hour*time.Duration(dineInTime.Hour()) + time.Minute*time.Duration(dineInTime.Minute()))
}

func GetReservationDB(c *gin.Context, resId int64) (models.Reservation, error) {

	var reservation models.Reservation
	query := `SELECT order_id, number_of_persons, dine_in_time, created_at, formatted_dine_in_time 
	FROM reservations WHERE id = $1`
	err := DbConn.QueryRow(query, resId).Scan(&reservation.OrderID, &reservation.NumberOfPersons, &reservation.DineInTime, &reservation.CreatedAt, &reservation.FormattedDineInTime)
	if err != nil {
		return models.Reservation{}, err
	}

	return reservation, nil
}
