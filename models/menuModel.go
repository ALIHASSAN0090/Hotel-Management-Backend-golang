package models

import "time"

type Menu struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Category  string    `json:"category" gorm:"type:varchar(100);not null"`
	StartDate time.Time `json:"start_date" gorm:"type:date;not null"`
	EndDate   time.Time `json:"end_date" gorm:"type:date;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	MenuID    int64     `json:"menu_id" gorm:"uniqueIndex;not null"`
}

type CreateMenu struct {
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Category string `json:"category" gorm:"type:varchar(100);not null"`
}

type UpdateMenu struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Category  string    `json:"category" gorm:"type:varchar(100);not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type FoodItem struct {
	ID     int64   `json:"id" db:"id"`
	Name   string  `json:"name" db:"name"`
	Price  float64 `json:"price" db:"price"`
	MenuID int64   `json:"menu_id" db:"menu_id"`
}

type MenuItems struct {
	ID        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Category  string     `json:"category" db:"category"`
	Foods     []FoodItem `json:"foods" db:"foods"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}
