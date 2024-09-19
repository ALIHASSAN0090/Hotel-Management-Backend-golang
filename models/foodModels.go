package models

import "time"

type Food struct {
	ID        int64     `json:"id" db:"primaryKey;autoIncrement"`
	Name      string    `json:"name" db:"type:varchar(100);not null"`
	Price     float64   `json:"price" db:"type:decimal(10,2);not null"`
	CreatedAt time.Time `json:"created_at" db:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" db:"autoUpdateTime"`
	MenuID    int64     `json:"menu_id" db:"type:varchar(100);not null"`
}
