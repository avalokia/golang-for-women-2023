package models

import "time"

type Product struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Variants  []Variants
	CreatedAt time.Time
	UpdatedAt time.Time
}
