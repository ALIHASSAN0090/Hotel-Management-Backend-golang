package models

import "time"

type Order struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderDate time.Time `json:"order_date" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	OrderID   int64     `json:"order_id" gorm:"uniqueIndex;not null"`
	TableID   int64     `json:"table_id" gorm:"not null"`
}
