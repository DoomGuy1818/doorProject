package psqlRepository

import (
	"doorProject/internal/domain/models"

	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (c ClientRepository) Create(client *models.Client) error {
	if err := c.db.Create(client).Error; err != nil {
		return err
	}

	return nil
}
