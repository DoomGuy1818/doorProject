package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	WorkerID  uint      `gorm:"not null" json:"worker_id" validate:"required"`
	ClientID  uint      `gorm:"not null" json:"client_id" validate:"required"`
	ServiceID uint      `gorm:"not null" json:"service_id" validate:"required"`
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
