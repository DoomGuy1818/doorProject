package models

import (
	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Name         string `json:"name" validate:"required"`
	Login        string `json:"login" validate:"required"`
	Password     string `json:"password" validate:"required"`
	IsActive     bool   `json:"is_active" default:"false"`
	WorkDays     []WorkerCalendar
	Appointments []Appointment
}
