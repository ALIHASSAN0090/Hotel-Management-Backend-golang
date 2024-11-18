package controllers

import (
	"golang-hotel-management/middleware"
	"golang-hotel-management/models"
	repo "golang-hotel-management/repositries/controllers_repo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repo repo.UserRepository
}

func NewUserController(repo repo.UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (uc *UserController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := uc.repo.GetUsersDB(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Message: "Users Fetched",
			Status:  200,
			Data:    data,
		})
	}
}

func (uc *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}
		UserData, err := uc.repo.GetUserDB(c, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			Message: "User Fetched",
			Status:  200,
			Data:    UserData,
		})
	}
}

func (uc *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User models.Login
		log.Printf("User struct initialized: %+v", User)

		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		log.Printf("Received login data: %+v", User)

		storedUser, err := uc.repo.GetUserByEmailDB(c, User.Email)
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
			log.Printf("Error generating token: %v", err)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Token generation failed"})
			return
		}

		log.Printf("Token generated successfully for user ID: %d", storedUser.ID)

		c.JSON(http.StatusOK, models.Response{
			Message: "Login Successful",
			Status:  200,
			Data:    gin.H{"token": token},
		})
	}
}

func (uc *UserController) Signup() gin.HandlerFunc {
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
				c.Error(err)
				return
			}

			if match, _ := VerifyPassword(getHash, User.PasswordHash); match {
				User.PasswordHash = getHash
				break
			}
		}

		log.Printf("Hashed password: %s", User.PasswordHash)

		createdUser, err := uc.repo.CreateUser(c, User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "User creation failed"})
			c.Error(err)
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

func (uc *UserController) CheckHeader() gin.HandlerFunc {
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
