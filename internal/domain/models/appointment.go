package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	WorkerId  uint      `json:"worker_id" validate:"required"`
	ClientId  uint      `json:"client_id" validate:"required"`
	ServiceId uint      `gorm:"not null" json:"service_id" validate:"required"`
	Date      time.Time `gorm:"type:date;not null" json:"date"`
	StartTime time.Time `gorm:"type:time;not null" json:"start_time"`
	EndTime   time.Time `gorm:"type:time;not null" json:"end_time"`
	// Статусы записи:
	// - "" (пусто)    - запись активна
	// - "skip"        - запись была пропущена
	// - "served"      - клиент обслужен
	// - "canceled"    - запись отменена
	Status string `json:"status"`
}
