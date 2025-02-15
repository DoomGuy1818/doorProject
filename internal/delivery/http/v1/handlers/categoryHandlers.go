package handlers

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: *s,
	}
}

func (h *CategoryHandler) CreateCategory(ctx echo.Context) error {
	c := new(models.Category)
	if err := ctx.Bind(c); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateCategory(c); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, c)
}
