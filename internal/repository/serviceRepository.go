package repository

import "doorProject/internal/domain/models"

type ServiceRepository interface {
	Create(service *models.Service) error
	FindServiceByIdAndWorker(serviceID uint, workerID uint) (*models.Service, error)
}
