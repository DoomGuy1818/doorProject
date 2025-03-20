package psqlRepository

import (
	"doorProject/internal/domain/models"
	"time"

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

func (wc *WorkerCalendarRepository) Create(workerCalendar *models.WorkerCalendar) error {
	if err := wc.database.Create(workerCalendar).Error; err != nil {
		return err
	}

	return nil
}

func (wc *WorkerCalendarRepository) FindCalendarByDateAndWorkerID(
	date time.Time,
	workerID uint,
) (*models.WorkerCalendar, error) {
	var workerCalendar *models.WorkerCalendar

	err := wc.database.Where(
		models.WorkerCalendar{
			Day:      date,
			WorkerID: workerID,
		},
	).First(&workerCalendar).Error
	if err != nil {
		return nil, err
	}
	return workerCalendar, nil
}
