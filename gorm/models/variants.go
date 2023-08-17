package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Variants struct {
	ID          uint `gorm:"primaryKey"`
	VariantName string
	Quantity    int
	ProductID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variants) BeforeCreate(tx *gorm.DB) (err error) {
	if (v.Quantity) < 0 {
		err = errors.New("Quantity cannot be less than 0.")
	}

	if len(v.VariantName) < 4 {
		err = errors.New("Variant name is too short. Minimum name is 4")
	}
	return
}
