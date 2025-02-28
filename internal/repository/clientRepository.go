package repository

import "doorProject/internal/domain/models"

type ClientRepository interface {
	Create(client *models.Client) error
}
