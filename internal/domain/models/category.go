package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title    string    `json:"name"`
	Products []Product `json:"products"`
}
