package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkerCalendar struct {
	gorm.Model
	Day       time.Time `gorm:"type:date; not null'" json:"day"`
	WorkStart time.Time `gorm:"type:date; not null'" json:"work_start"`
	WorkEnd   time.Time `gorm:"type:date; not null'" json:"work_end"`
	ServiceId uint      `json:"service_id"`
	WorkerId  uint      `json:"worker_id"`
}
