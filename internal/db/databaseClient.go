package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseClient struct {
	dsn string
}

func NewDatabaseClient(dsn string) *DatabaseClient {
	return &DatabaseClient{dsn: dsn}
}

func (d DatabaseClient) GetDBClient() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(d.dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err
}
