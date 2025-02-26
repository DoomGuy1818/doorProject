package models

import (
	"time"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Name      string    `json:"name" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Login     string    `json:"login" validate:"required"`
	WorkDay   time.Time `gorm:"type:date;not null" json:"work_day" validate:"required"`
	WorkStart time.Time `json:"work_start" validate:"required"`
	WorkEnd   time.Time `json:"work_end" validate:"required"`
}
