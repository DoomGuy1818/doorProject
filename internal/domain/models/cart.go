package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ClientID uint      `json:"client_id"`
	Bill     float64   `json:"bill" validate:"required"`
	Products []Product `gorm:"many2many:product_carts"`
}
