package configs

import (
	"doorProject/internal/domain/models"
	"log"

	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Database *gorm.DB
}

func (database DatabaseConfig) SetupDb(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Cart{},
		&models.Worker{},
		&models.Client{},
		&models.Service{},
		&models.Color{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection established successfully")
	database.Database = db
}
