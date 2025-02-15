package handlers

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/service"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service   service.CategoryService
	validator *validator.Validate
}

func NewCategoryHandler(s *service.CategoryService, v *validator.Validate) *CategoryHandler {
	return &CategoryHandler{
		service:   *s,
		validator: v,
	}
}

func (h *CategoryHandler) CreateCategory(ctx echo.Context) error {
	c := new(models.Category)
	if err := ctx.Bind(c); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(c); err != nil {
		if errors.As(err, &validator.ValidationErrors{}) {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
	}

	if err := h.service.CreateCategory(c); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, c)
}
