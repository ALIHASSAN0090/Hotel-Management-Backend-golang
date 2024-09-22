package models

import "time"

type Invoice struct {
	ID            int64     `json:"id" db:"id"`
	OrderID       int64     `json:"order_id" db:"order_id"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	PaymentStatus string    `json:"payment_status" db:"payment_status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
