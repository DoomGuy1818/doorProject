package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type AuthRoutes struct {
	authHandler *handlers.AuthHandlers
}

func NewAuthRoutes(authHandler *handlers.AuthHandlers) *AuthRoutes {
	return &AuthRoutes{
		authHandler: authHandler,
	}
}

func (r *AuthRoutes) SighIn(e *echo.Echo) {
	e.POST("auth/signin", r.authHandler.SignIn)
}
