package controller_repo

import "github.com/gin-gonic/gin"

type UserController interface {
	GetUsers() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	Login() gin.HandlerFunc
	Signup() gin.HandlerFunc
	HashPassword(password string) (string, error)
	VerifyPassword(userPassword string, providedPassword string) (bool, string)
	CheckHeader() gin.HandlerFunc
}
