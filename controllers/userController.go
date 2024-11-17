package controllers

import (
	"golang-hotel-management/database"
	"golang-hotel-management/middleware"
	"golang-hotel-management/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		data, err := database.GetUsersDB(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "User Fetched",
			Status:  200,
			Data:    data,
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var User models.Login
		log.Printf("User struct initialized: %+v", User)

		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		log.Printf("Received login data: %+v", User)

		storedUser, err := database.GetUserByEmailDB(c, User.Email)
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Email not found"})
			return
		}

		if match, _ := VerifyPassword(storedUser.PasswordHash, User.PasswordHash); !match {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid credentials"})
			return
		}

		token, err := middleware.GenerateToken(uint(storedUser.ID), storedUser.Email, storedUser.Username, storedUser.Role)
		if err != nil {
			log.Printf("Error generating token: %v", err) // Log the error
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Token generation failed"})
			return
		}

		log.Printf("Token generated successfully for user ID: %d", storedUser.ID) // Log successful token generation

		c.JSON(http.StatusOK, models.Response{
			Message: "Login Successful",
			Status:  200,
			Data:    gin.H{"token": token},
		})
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

		log.Printf("Hashed password: %s", User.PasswordHash)

		createdUser, err := database.CreateUser(c, User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "User creation failed"})
			c.Error(err) // Log the error
			return
		}

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

func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		email, _ := c.Get("email")
		username, _ := c.Get("username")
		role, _ := c.Get("role")

		c.JSON(http.StatusOK, gin.H{
			"userID":   userID,
			"email":    email,
			"username": username,
			"role":     role,
		})
	}
}
