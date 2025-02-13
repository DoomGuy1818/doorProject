package models

import "gorm.io/gorm"

type Color struct {
	gorm.Model
	Color    string    `json:"color" validate:"required"`
	Products []Product `json:"products"`
}
