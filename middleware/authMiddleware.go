package middleware

import (
	"fmt"
	"golang-hotel-management/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID uint, email, username, role string) (string, error) {
	claims := &models.CustomClaims{
		ID:       fmt.Sprintf("%d", userID),
		Email:    email,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "golang-hotel-management",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &models.CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userID := claims.ID
		email := claims.Email
		username := claims.Username
		role := claims.Role
		fmt.Printf("Extracted Claims - ID: %s, Email: %s, Username: %s , role : %s\n", userID, email, username, role)

		c.Set("userID", userID)
		c.Set("email", email)
		c.Set("username", username)
		c.Set("role", role)
		c.Next()
	}
}
