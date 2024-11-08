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

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		var NewMenu models.CreateMenu

		if err := c.ShouldBindJSON(&NewMenu); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		}

		createdMenu, err := database.CreateMenuDB(c, NewMenu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to Create Menu Item"})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			Message: "Menu Created Succesfully",
			Status:  200,
			Data:    createdMenu,
		})
	}
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var UpdateMenu models.UpdateMenu

		if err := c.ShouldBindJSON(&UpdateMenu); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		}

		data, err := database.UpdateMenuDB(c, UpdateMenu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to Create Menu Item"})
		}
		c.JSON(http.StatusOK, models.Response{
			Message: "Menu Created Succesfully",
			Status:  200,
			Data:    data,
		})
	}
}
