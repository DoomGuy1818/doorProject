package handlers

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CartHandlers struct {
	service   *service.CartService
	validator *validator.Validate
}

func NewCartHandlers(service *service.CartService, validator *validator.Validate) *CartHandlers {
	return &CartHandlers{
		service:   service,
		validator: validator,
	}
}

func (crt *CartHandlers) CreateCart(ctx echo.Context) error {
	cartDto := new(dto.CartDto)

	if err := ctx.Bind(cartDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := crt.validator.Struct(cartDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	form, err := crt.service.CreateCart(cartDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, form)
}
