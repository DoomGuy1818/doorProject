package repository

import (
	"doorProject/internal/domain/models"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(id int, product models.Product) error
	DeleteProduct(id int) error
	GetProductById(id int) error
	GetProducts() error
}
