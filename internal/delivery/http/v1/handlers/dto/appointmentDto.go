package dto

import "time"

type AppointmentDto struct {
	MasterId  uint      `json:"master_id"`
	ClientId  uint      `json:"client_id" validate:"required"`
	ServiceId uint      `json:"service_id" validate:"required"`
	Date      time.Time `gorm:"type:date;not null" json:"date"`
	StartTime time.Time `gorm:"type:date;not null" json:"start_time"`
	EndTime   time.Time `gorm:"type:date;not null" json:"end_time"`
	Status    string    `json:"status" validate:"required"`
}
