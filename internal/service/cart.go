package service

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
)

type CartService struct {
	repository repository.CartRepository
}

func NewCartService(repository repository.CartRepository) *CartService {
	return &CartService{
		repository: repository,
	}
}

func (crt *CartService) CreateCart(dto *dto.CartDto) (*models.Cart, error) {
	cart := &models.Cart{
		ClientId: dto.ClientId,
		Bill:     dto.Bill,
		Products: dto.Products,
	}

	if err := crt.repository.Create(cart); err != nil {
		return nil, err
	}

	return cart, nil
}
