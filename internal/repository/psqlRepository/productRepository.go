package psqlRepository

import (
	"doorProject/internal/domain/models"
	"log"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) CreateProduct(product *models.Product) error {
	if err := p.db.Create(product).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *ProductRepository) UpdateProduct(id int, product models.Product) error {
	var existingProduct models.Product

	// Находим существующий продукт
	if err := p.db.First(&existingProduct, id).Error; err != nil {
		log.Println(err)
		return err
	}

	// Обновляем поля
	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Amount = product.Amount
	existingProduct.UpdatedAt = product.UpdatedAt
	existingProduct.IsActive = product.IsActive
	existingProduct.Price = product.Price

	// Сохраняем изменения
	if err := p.db.Save(&existingProduct).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DeleteProduct удаляет продукт из базы данных
func (p *ProductRepository) DeleteProduct(id int) error {
	if err := p.db.Delete(&models.Product{}, id).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetProductById получает продукт по ID
func (p *ProductRepository) GetProductById(id int) error {
	var product models.Product
	if err := p.db.First(&product, id).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetProducts получает все продукты
func (p *ProductRepository) GetProducts() error {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
