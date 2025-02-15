package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type CategoryRoute struct {
	handler *handlers.CategoryHandler
}

func NewCategoryRoute(handler *handlers.CategoryHandler) *CategoryRoute {
	return &CategoryRoute{
		handler: handler,
	}
}

func (r CategoryRoute) CreateCategory(echo *echo.Echo) {
	echo.POST("/categories", r.handler.CreateCategory)
}
