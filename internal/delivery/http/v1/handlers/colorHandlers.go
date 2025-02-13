package handlers

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type СolorHandler struct {
	service service.ColorService
}

func NewColorHandler(colorService *service.ColorService) *СolorHandler {
	return &СolorHandler{
		service: *colorService,
	}
}

func (h *СolorHandler) CreateColor(ctx echo.Context) error {
	p := new(models.Color)

	if err := ctx.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateColor(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, p)
}
