package controller_repo

import (
	"golang-hotel-management/models"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetOrders() gin.HandlerFunc
	GetOrder() gin.HandlerFunc
	CreateOrder() gin.HandlerFunc
	UpdateOrder() gin.HandlerFunc
	PrepareAndSendReservationEmail(c *gin.Context, createOrder models.CombinedOrderReservation) error
	PrepareAndSendOrderEmail(c *gin.Context, order models.CombinedOrderReservation) error
}
