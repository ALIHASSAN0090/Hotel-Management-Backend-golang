package controllers

import (
	"golang-hotel-management/database"
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMenusWithFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := database.GetAllMenusWithFoodsDB(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Fetched All Menus",
			Status:  200,
			Data:    data,
		})

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
