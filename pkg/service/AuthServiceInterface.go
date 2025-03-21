package service

import (
	"doorProject/internal/domain/models"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	GenerateTokenAndSetCookie(user *models.Worker, c echo.Context) error
	TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}
