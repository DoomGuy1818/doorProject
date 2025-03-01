package psqlRepository

import (
	"doorProject/internal/domain/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (crt CartRepository) Create(cart *models.Cart) error {
	if err := crt.db.Create(cart).Error; err != nil {
		return err
	}

	return nil
}
