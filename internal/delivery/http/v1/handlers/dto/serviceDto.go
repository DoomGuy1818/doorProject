package dto

import "time"

type ServiceDto struct {
	Name     string        `json:"name" validate:"required"`
	IsActive bool          `json:"is_active" default:"true"`
	Price    float64       `json:"price" validate:"required"`
	Duration time.Duration `json:"duration" validate:"required"`
	WorkerId uint          `json:"worker_id" validate:"required"`
}
