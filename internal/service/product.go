package service

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(db repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: db,
	}
}

func (p ProductService) CreateProduct(dto *models.Product) error {
	result := p.productRepository.CreateProduct(dto)

	if result != nil {
		return result
	}

	return nil
}

func updateProduct(id string) error {
	return nil
}

func deleteProduct(id string) error {
	return nil
}

func getProduct(id string) error {
	return nil
}
