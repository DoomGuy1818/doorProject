package psqlRepository

import (
	"doorProject/internal/domain/models"

	"gorm.io/gorm"
)

type WorkerRepository struct {
	db *gorm.DB
}

func NewWorkerRepository(db *gorm.DB) *WorkerRepository {
	return &WorkerRepository{db: db}
}

func (w WorkerRepository) Create(worker *models.Worker) error {
	if err := w.db.Create(worker).Error; err != nil {
		return err
	}

	return nil
}
