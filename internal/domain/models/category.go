package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title    string    `json:"name" validate:"required,min=2,max=20"`
	Products []Product `json:"products"`
}
