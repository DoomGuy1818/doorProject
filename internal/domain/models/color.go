package models

import "gorm.io/gorm"

type Color struct {
	gorm.Model
	Color    string    `json:"string"`
	Products []Product `json:"products"`
}
