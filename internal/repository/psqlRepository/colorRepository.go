package psqlRepository

import (
	"doorProject/internal/domain/models"

	"gorm.io/gorm"
)

type ColorRepository struct {
	db *gorm.DB
}

func NewColorRepository(db *gorm.DB) *ColorRepository {
	return &ColorRepository{db: db}
}

func (c ColorRepository) CreateColor(color *models.Color) error {
	if err := c.db.Create(&color).Error; err != nil {
		return err
	}

	return nil
}
