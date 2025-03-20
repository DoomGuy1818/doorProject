package repository

import (
	"doorProject/internal/domain/models"
	"time"
)

type WorkerCalendarRepository interface {
	Create(workerCalendar *models.WorkerCalendar) error
	FindCalendarByDateAndWorkerID(date time.Time, workerID uint) (*models.WorkerCalendar, error)
}
