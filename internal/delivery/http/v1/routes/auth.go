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

func (r *AuthRoutes) AuthRoutes(e *echo.Echo) {
	e.POST("auth/signin", r.authHandler.SignIn)
	//e.POST("auth/signout", r.authHandler.SignOut)
	//e.POST("auth/register", r.authHandler.Register)
	//e.POST("auth/refresh", r.authHandler.Refresh)
}
