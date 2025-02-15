package handlers

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/service"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type СolorHandler struct {
	service   service.ColorService
	validator *validator.Validate
}

func NewColorHandler(colorService *service.ColorService, v *validator.Validate) *СolorHandler {
	return &СolorHandler{
		service:   *colorService,
		validator: v,
	}
}

func (h *СolorHandler) CreateColor(ctx echo.Context) error {
	c := new(models.Color)

	if err := ctx.Bind(c); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(c); err != nil {
		if errors.As(err, &validator.ValidationErrors{}) {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
	}

	if err := h.service.CreateColor(c); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, c)
}
