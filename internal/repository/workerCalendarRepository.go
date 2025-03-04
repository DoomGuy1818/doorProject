package repository

import (
	"doorProject/internal/domain/models"
	"time"
)

type WorkerCalendarRepository interface {
	Create(workerCalendar *models.WorkerCalendar) error
	GetByWorkerAndDate(workerId uint, date time.Time) (*models.WorkerCalendar, error)
}
