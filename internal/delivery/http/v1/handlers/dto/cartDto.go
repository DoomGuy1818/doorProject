package dto

import "doorProject/internal/domain/models"

type CartDto struct {
	ClientId uint             `json:"client_id" validate:"required"`
	Bill     float64          `json:"bill" default:"0.0"`
	Products []models.Product `json:"products" default:"[]"`
}
