package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type ClientRoutes struct {
	handlers *handlers.ClientHandlers
}

func NewClientRoutes(handler *handlers.ClientHandlers) *ClientRoutes {
	return &ClientRoutes{
		handlers: handler,
	}
}

func (r *ClientRoutes) CreateClient(e *echo.Echo) {
	e.POST("/clients", r.handlers.Create)
}
