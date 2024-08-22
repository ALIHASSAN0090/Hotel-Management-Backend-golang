package models

import "time"

type Invoice struct {
	ID             int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	InvoiceID      string    `json:"invoice_id" gorm:"type:varchar(100);not null;unique"`
	OrderID        int64     `json:"order_id" gorm:"not null"`
	PaymentMethod  string    `json:"payment_method" gorm:"type:varchar(50);not null"`
	PaymentStatus  string    `json:"payment_status" gorm:"type:varchar(50);not null"`
	PaymentDueDate time.Time `json:"payment_due_date" gorm:"type:date;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
