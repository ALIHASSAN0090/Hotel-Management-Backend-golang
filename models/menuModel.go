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
