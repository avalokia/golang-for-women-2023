package models

import "time"

type Score struct {
	ID              uint   `gorm:"primaryKey"`
	AssignmentTitle string `gorm:"not null;type:varchar" json:"assignment_title"`
	Description     string `gorm:"not null;type:varchar" json:"description"`
	Score           int    `gorm:"not null" json:"score"`
	StudentID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
