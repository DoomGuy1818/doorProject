package repository

import "doorProject/internal/domain/models"

type WorkerCalendarRepository interface {
	Create(workerCalendar *models.WorkerCalendar) error
}
