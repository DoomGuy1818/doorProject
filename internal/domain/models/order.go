package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Title           string    `json:"title" validate:"required"`
	IsCanSendNotify bool      `json:"is_can_send_notify" default:"false"`
	DateStart       time.Time `json:"date_start" validate:"required"`
	DateEnd         time.Time `json:"date_end" validate:"required"`
	Workers         []Worker  `gorm:"many2many:order_workers"`
}
