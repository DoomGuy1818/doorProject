package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type ColorRoutes struct {
	handlers *handlers.СolorHandler
}

func NewColorRoutes(handler *handlers.СolorHandler) *ColorRoutes {
	return &ColorRoutes{
		handlers: handler,
	}
}

func (c ColorRoutes) CreateColor(echo *echo.Echo) {
	echo.POST("/colors", c.handlers.CreateColor)
}
