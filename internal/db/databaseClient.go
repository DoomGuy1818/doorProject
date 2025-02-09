package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDsn() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	return os.Getenv("DB_URL")
}

func GetDBClient() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDsn()), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err
}
