package models

import "time"

type WorkingTime struct {
	Date        time.Time `gorm:"primaryKey;type:date" json:"date"`
	StartTime   time.Time `gorm:"type:time" json:"start_time"`
	EndTime     time.Time `gorm:"type:time" json:"end_time"`
	MaxHours    int       `gorm:"default:8" json:"max_hours"`    // В часах
	BookedHours int       `gorm:"default:0" json:"booked_hours"` // В минутах
}
