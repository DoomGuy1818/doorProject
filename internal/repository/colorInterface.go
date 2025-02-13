package repository

import "doorProject/internal/domain/models"

type ColorRepository interface {
	CreateColor(color *models.Color) error
}
