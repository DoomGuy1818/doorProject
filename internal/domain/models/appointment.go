package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	ClientID  uint      `json:"client_id"` // Ссылка на таблицу клиентов
	ServiceID uint      `gorm:"not null" json:"service_id"`
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
