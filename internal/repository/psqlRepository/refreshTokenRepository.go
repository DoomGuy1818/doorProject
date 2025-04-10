package psqlRepository

import (
	"doorProject/internal/domain/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) CreateRefreshToken(refreshToken *models.UserToken) error {
	if err := r.db.Create(refreshToken).Error; err != nil {
		return err
	}

	return nil
}

func (r *RefreshTokenRepository) RemoveRefreshToken(ID int) error {
	if err := r.db.Delete(&models.UserToken{}, ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *RefreshTokenRepository) FindRefreshTokenByToken(token string) (*models.UserToken, error) {
	var refreshToken models.UserToken
	if err := r.db.Where("token_hash = ?", token).First(&refreshToken).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *RefreshTokenRepository) InvalidateRefreshToken(token *models.UserToken) error {
	token.IsValid = false
	err := r.db.Save(&token)
	if err != nil {
		return err.Error
	}

	return nil
}
