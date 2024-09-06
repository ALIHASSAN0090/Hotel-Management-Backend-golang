package models

import "time"

type Note struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Text      string     `json:"text" gorm:"type:text;not null"`
	Title     string     `json:"title" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"` // Soft delete field
	NoteID    int64      `json:"note_id" gorm:"unique;not null"`
}
