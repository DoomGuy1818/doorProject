package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type ProductRoute struct {
	handler *handlers.ProductHandler
}

func NewProductRoute(hand *handlers.ProductHandler) *ProductRoute {
	return &ProductRoute{
		handler: hand,
	}
}

func (p *ProductRoute) CreateProduct(echo *echo.Echo) {
	echo.POST("/product", p.handler.CreateProduct)
}
