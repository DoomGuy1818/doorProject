package repository

import "doorProject/internal/domain/models"

type WorkerRepositoryInterface interface {
	Create(worker *models.Worker) error
	FindUserByUsername(username string) (*models.Worker, error)
}
