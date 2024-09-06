package models

import "time"

type Table struct {
	ID             int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	NumberOfGuests int       `json:"number_of_guests" gorm:"not null"`
	TableNumber    int       `json:"table_number" gorm:"uniqueIndex;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TableID        int64     `json:"table_id" gorm:"uniqueIndex;not null"`
}
