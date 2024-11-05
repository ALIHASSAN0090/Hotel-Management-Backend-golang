package models

import (
	"time"
)

type Order struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderDate time.Time `json:"order_date" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	OrderID   int64     `json:"order_id" gorm:"uniqueIndex;not null"`
	TableID   int64     `json:"table_id" gorm:"not null"`
}

type AllOrders struct {
	OrderID    int64       `json:"order_id" db:"order_id"`
	TotalPrice float64     `json:"total_price" db:"total_price"`
	Status     string      `json:"status" db:"status"`
	FoodItems  []FoodItems `json:"food_items" db:"food_items"`
}

type FoodItems struct {
	FoodItemID   int64   `json:"food_item_id" db:"food_item_id"`
	Name         string  `json:"name" db:"name"`
	PricePerUnit float64 `json:"price_per_unit" db:"price_per_unit"`
	FoodCategory string  `json:"food_category" db:"food_category"`
	Category     string  `json:"category" db:"category"`
}

type CreateOrder struct {
	UserID        int64    `json:"user_id" db:"user_id"`
	FoodItems_IDs []FoodID `json:"food_items" db:"food_items"`
}

type FoodID struct {
	ID int64 `json:"id" db:"id"`
}
