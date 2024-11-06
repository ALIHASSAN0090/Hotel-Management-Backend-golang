package controllers

import (
	"golang-hotel-management/database"
	"golang-hotel-management/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

		var User models.User
		log.Printf("User struct initialized: %+v", User)

		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		log.Printf("Received user data: %+v", User)

		for {
			getHash, err := HashPassword(User.PasswordHash)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Password hashing failed"})
				c.Error(err) // Log the error
				return
			}

			if match, _ := VerifyPassword(getHash, User.PasswordHash); match {
				User.PasswordHash = getHash
				break
			}
		}

		// Debugging: Log the hashed password
		log.Printf("Hashed password: %s", User.PasswordHash)

		createdUser, err := database.CreateUser(c, User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "User creation failed"})
			c.Error(err) // Log the error
			return
		}

		// Debugging: Log the created user data
		log.Printf("Created user data: %+v", createdUser)

		c.JSON(http.StatusOK, models.Response{
			Message: "User Created",
			Status:  200,
			Data:    createdUser,
		})
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}
	log.Printf("Password hashed successfully")
	return string(hashedPassword), nil
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		log.Printf("Password verification failed: %v", err)
		return false, "Password does not match"
	}
	log.Printf("Password verified successfully")
	return true, "Password matched"
}
