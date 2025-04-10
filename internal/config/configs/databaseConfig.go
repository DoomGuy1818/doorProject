package configs

import (
	"doorProject/internal/domain/models"
	"log"

	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Database *gorm.DB
}

func NewDatabaseConfig(db *gorm.DB, err error) *DatabaseConfig {
	if err != nil {
		log.Fatal(err)
	}

	return &DatabaseConfig{
		Database: db,
	}
}

func (database DatabaseConfig) SetupDb() {
	err := database.Database.AutoMigrate(
		&models.Cart{},
		&models.Worker{},
		&models.Client{},
		&models.Service{},
		&models.Color{},
		&models.Category{},
		&models.Product{},
		&models.Appointment{},
		&models.WorkerCalendar{},
		&models.UserToken{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection established successfully")
}
