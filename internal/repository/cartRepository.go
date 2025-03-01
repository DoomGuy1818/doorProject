package repository

import "doorProject/internal/domain/models"

type CartRepository interface {
	Create(cart *models.Cart) error
}
