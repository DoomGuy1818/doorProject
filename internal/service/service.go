package service

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
)

type Service struct {
	repository repository.ServiceRepository
}

func NewService(repository repository.ServiceRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) CreateService(dto *dto.ServiceDto) (*models.Service, error) {
	service := &models.Service{
		Name:     dto.Name,
		IsActive: dto.IsActive,
		Price:    dto.Price,
		Duration: dto.Duration,
		WorkerId: dto.WorkerId,
	}

	if err := s.repository.Create(service); err != nil {
		return nil, err
	}

	return service, nil
}
