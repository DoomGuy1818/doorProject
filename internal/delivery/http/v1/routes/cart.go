package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type CartRoutes struct {
	handler *handlers.CartHandlers
}

func NewCartRoutes(handler *handlers.CartHandlers) *CartRoutes {
	return &CartRoutes{
		handler: handler,
	}
}

func (r *CartRoutes) CreateCart(echo *echo.Echo) {
	echo.POST("/cart", r.handler.CreateCart)
}
