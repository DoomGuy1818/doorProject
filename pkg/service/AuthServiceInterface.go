package service

import (
	"doorProject/internal/domain/models"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	GenerateTokenAndSetCookie(user *models.Worker, c echo.Context) (string, error)
	TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	JWTErrorChecker(c echo.Context, err error) error
}
