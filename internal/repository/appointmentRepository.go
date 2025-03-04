package repository

import "doorProject/internal/domain/models"

type AppointmentRepository interface {
	Create(appointment *models.Appointment) error
}
