package psqlRepository

import (
	"doorProject/internal/domain/models"

	"gorm.io/gorm"
)

type WorkerCalendarRepository struct {
	database *gorm.DB
}

func NewWorkerCalendarRepository(db *gorm.DB) *WorkerCalendarRepository {
	return &WorkerCalendarRepository{
		database: db,
	}
}

func (wc WorkerCalendarRepository) Create(workerCalendar *models.WorkerCalendar) error {
	if err := wc.database.Create(workerCalendar).Error; err != nil {
		return err
	}

	return nil
}
