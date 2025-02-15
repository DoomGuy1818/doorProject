package psqlRepository

import (
	"doorProject/internal/domain/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (c CategoryRepository) CreateCategory(category *models.Category) error {
	if err := c.db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}
