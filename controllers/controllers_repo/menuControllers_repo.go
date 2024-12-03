package controller_repo

import "github.com/gin-gonic/gin"

type MenuController interface {
	GetAllMenusWithFoods() gin.HandlerFunc
	CreateMenu() gin.HandlerFunc
	UpdateMenu() gin.HandlerFunc
}
