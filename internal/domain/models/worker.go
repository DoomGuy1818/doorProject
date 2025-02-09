package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Name     string  `json:"name" validate:"required"`
	Password string  `json:"password" validate:"required"`
	Login    string  `json:"login" validate:"required"`
	Orders   []Order `gorm:"many2many:worker_orders"`
}
