package models

import "time"

type OrderItem struct {
	ID          int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Quantity    string     `json:"quantity" gorm:"type:varchar(20);not null"` // Categorical value
	UnitPrice   float64    `json:"unit_price" gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"` // Soft delete field
	FoodID      string     `json:"food_id" gorm:"not null"`
	OrderItemID string     `json:"order_item_id" gorm:"unique;not null"`
	OrderID     string     `json:"order_id" gorm:"not null"`
}
