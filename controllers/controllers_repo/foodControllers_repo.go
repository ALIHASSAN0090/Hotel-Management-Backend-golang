package controller_repo

import "github.com/gin-gonic/gin"

type FoodController interface {
	GetAllFoods() gin.HandlerFunc
	GetFood() gin.HandlerFunc
	CreateFood() gin.HandlerFunc
	UpdateFood() gin.HandlerFunc
}
