package psqlRepository

import (
	"doorProject/internal/domain/models"
	"time"

	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{
		db: db,
	}
}

func (a *AppointmentRepository) Create(appointment *models.Appointment) error {
	if err := a.db.Create(appointment).Error; err != nil {
		return err
	}

	return nil
}

func (a *AppointmentRepository) FindAppointmentsByDay(date time.Time) ([]models.Appointment, error) {
	var appointments []models.Appointment
	if err := a.db.Where(
		models.Appointment{
			Date: date,
		},
	).Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}
