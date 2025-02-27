package models

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name     string    `json:"name" validate:"required"`
	IsActive bool      `json:"is_active" default:"true"`
	Price    float64   `json:"price" validate:"required"`
	Duration time.Time `json:"duration" validate:"required"`
	WorkerId int       `gorm:"not null" json:"worker_id" validate:"required"`
}
