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
