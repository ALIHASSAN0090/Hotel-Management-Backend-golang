package database_repo

import (
	"golang-hotel-management/models"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderRepository interface {
	GetAllOrdersDB(c *gin.Context) ([]models.AllOrders, error)
	GetOrderDB(c *gin.Context, incmoingOrderId int64) (*models.AllOrders, error)
	CreateOrderDB(c *gin.Context, createOrder models.CombinedOrderReservation) (int64, int64, error)
	GetOrderFoodsDB(foodids []models.FoodID) (string, error)
	GetReservationDB(c *gin.Context, resId int64) (models.Reservation, error)
	GetOrderByUserIdDB(userid any, c *gin.Context) (string, error)
	GetTotalPrice(foodids []models.FoodID) (float64, error)
	UpdateOrderDB(c *gin.Context, updateId int64) (models.AllOrders, error)
	CalculateDineInDateTime(dineInDate, dineInTime time.Time) time.Time
}
