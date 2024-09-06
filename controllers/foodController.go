package controllers

import (
	"context"
	"golang-hotel-management/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		food_id := c.Param("food_id")

		var food models.Food
		

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Round(num float64) int {

}

func Fixed(num float64, preceision int) float64 {

}
