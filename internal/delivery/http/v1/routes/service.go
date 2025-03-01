package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type ServiceRoutes struct {
	handler *handlers.ServiceHandlers
}

func NewServiceRoutes(handler *handlers.ServiceHandlers) *ServiceRoutes {
	return &ServiceRoutes{
		handler: handler,
	}
}

func (r *ServiceRoutes) CreateService(echo *echo.Echo) {
	echo.POST("/services", r.handler.CreateService)
}
