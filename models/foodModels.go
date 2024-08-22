package models

import "time"

type Food struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	FoodImage string    `json:"food_image" gorm:"type:varchar(255)"`
	FoodID    int64     `json:"food_id" gorm:"not null"`
	MenuID    int64     `json:"menu_id" gorm:"not null"`
}
