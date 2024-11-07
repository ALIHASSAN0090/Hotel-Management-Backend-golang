package models

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
)

type User struct {
	ID           int64      `json:"id" db:"id"`
	Username     string     `json:"username" db:"username"`
	Email        string     `json:"email" db:"email"`
	PasswordHash string     `json:"password_hash" db:"password_hash"`
	FirstName    string     `json:"first_name" db:"first_name"`
	LastName     string     `json:"last_name" db:"last_name"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	Role         string     `json:"role" db:"role"`
}

type CustomClaims struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Role     string `json:"role" db:"role"`
	jwt.StandardClaims
}

type Login struct {
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}
