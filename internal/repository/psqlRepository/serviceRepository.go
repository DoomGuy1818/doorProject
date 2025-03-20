package psqlRepository

import (
	"doorProject/internal/domain/models"

	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (s ServiceRepository) Create(service *models.Service) error {
	if err := s.db.Create(service).Error; err != nil {
		return err
	}
	return nil
}

func (s ServiceRepository) FindServiceByIdAndWorker(serviceID uint, workerID uint) (*models.Service, error) {
	var service models.Service
	if err := s.db.Where("id = ? AND worker_id = ?", serviceID, workerID).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}
