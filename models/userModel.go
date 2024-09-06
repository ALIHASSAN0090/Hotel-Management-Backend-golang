package models

import "time"

type User struct {
	ID           int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName    string     `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName     string     `json:"last_name" gorm:"type:varchar(100);not null"`
	Password     string     `json:"password" gorm:"type:varchar(255);not null"` // Consider hashing passwords
	Email        string     `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Avatar       string     `json:"avatar" gorm:"type:varchar(255)"`
	Phone        string     `json:"phone" gorm:"type:varchar(20)"`
	Token        string     `json:"token" gorm:"type:text"`
	RefreshToken string     `json:"refresh_token" gorm:"type:text"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" gorm:"index"` // Soft delete field
	UserID       int64      `json:"user_id" gorm:"unique;not null"`
}
