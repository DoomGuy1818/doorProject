package psqlRepository

import (
	"doorProject/internal/domain/models"

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
