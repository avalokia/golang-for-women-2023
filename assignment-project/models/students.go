package models

import "time"

type Student struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"not null;type:varchar" json:"name"`
	Age       int     `gorm:"not null" json:"age"`
	Scores    []Score `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
