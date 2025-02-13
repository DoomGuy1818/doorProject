package service

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
)

type ColorService struct {
	repository repository.ColorRepository
}

func NewColorService(repository repository.ColorRepository) *ColorService {
	return &ColorService{
		repository: repository,
	}
}

func (c ColorService) CreateColor(dto *models.Color) error {
	result := c.repository.CreateColor(dto)
	if result != nil {
		return result
	}
	return nil
}
