package service

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
)

type CategoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (c CategoryService) CreateCategory(dto *models.Category) error {
	result := c.repository.CreateCategory(dto)
	if result != nil {
		return result
	}

	return nil
}
