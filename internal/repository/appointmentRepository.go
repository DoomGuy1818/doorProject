package repository

import (
	"doorProject/internal/domain/models"
	"time"
)

type AppointmentRepository interface {
	Create(appointment *models.Appointment) error
	GetByWorkerAndDate(workerId uint, date time.Time) ([]models.Appointment, error)
}
