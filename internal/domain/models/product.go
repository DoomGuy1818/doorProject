package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required"` // Используйте float64 для цены
	Amount      float64 `json:"amount" validate:"required"`
	IsActive    bool    `json:"is_active" default:"true"`
	CategoryID  uint    `json:"category"`
	ColorID     uint    `json:"color"`
	Carts       []Cart  `gorm:"many2many:product_carts"`
}
