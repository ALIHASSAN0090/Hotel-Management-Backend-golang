package controllers

import (
	controller_repo "golang-hotel-management/controllers/controllers_repo"
	"golang-hotel-management/database/database_repo"
	"golang-hotel-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	Repo database_repo.MenuRepository
}

func NewMenuController(repo database_repo.MenuRepository) controller_repo.MenuController {
	return &MenuController{Repo: repo}
}

func (mc *MenuController) GetAllMenusWithFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := mc.Repo.GetAllMenusWithFoodsDB(c)
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

func (mc *MenuController) CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		var NewMenu models.CreateMenu

		if err := c.ShouldBindJSON(&NewMenu); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		}

		createdMenu, err := mc.Repo.CreateMenuDB(c, NewMenu)
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
func (mc *MenuController) UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var UpdateMenu models.UpdateMenu

		if err := c.ShouldBindJSON(&UpdateMenu); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		}

		data, err := mc.Repo.UpdateMenuDB(c, UpdateMenu)
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
