package repository

import (
	"doorProject/internal/domain/models"
	"time"
)

type AppointmentRepository interface {
	Create(appointment *models.Appointment) error
	FindAppointmentsByDay(date time.Time) ([]models.Appointment, error)
}
