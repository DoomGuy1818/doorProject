package handlers

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ClientHandlers struct {
	service   *service.Client
	validator *validator.Validate
}

func NewClientHandlers(service *service.Client, validator *validator.Validate) *ClientHandlers {
	return &ClientHandlers{
		service:   service,
		validator: validator,
	}
}

func (c *ClientHandlers) Create(ctx echo.Context) error {
	clientDto := new(dto.ClientDto)

	if err := ctx.Bind(clientDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.validator.Struct(clientDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	form, err := c.service.CreateClient(clientDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, form)
}
