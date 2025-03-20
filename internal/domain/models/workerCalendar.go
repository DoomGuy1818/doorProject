package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkerCalendar struct {
	gorm.Model
	Day       time.Time `gorm:"type:date; not null'" json:"day"`
	WorkStart time.Time `gorm:"type:time; not null'" json:"work_start"`
	WorkEnd   time.Time `gorm:"type:time; not null'" json:"work_end"`
	ServiceID uint      `json:"service_id"`
	WorkerID  uint      `json:"worker_id"`
}
