package repository

import "doorProject/internal/domain/models"

type RefreshTokenRepository interface {
	CreateRefreshToken(refreshToken *models.UserToken) error
	FindRefreshTokenByToken(token string) (*models.UserToken, error)
	InvalidateRefreshToken(token *models.UserToken) error
}
