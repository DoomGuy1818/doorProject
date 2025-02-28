package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type Service struct {
	handlers *handlers.ServiceHandlers
}

func NewService(handler *handlers.ServiceHandlers) *Service {
	return &Service{
		handlers: handler,
	}
}

func (r *Service) CreateService(echo *echo.Echo) {
	echo.POST("/services", r.handlers.CreateService)
}
