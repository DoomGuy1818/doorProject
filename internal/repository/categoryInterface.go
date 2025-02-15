package repository

import "doorProject/internal/domain/models"

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
}
