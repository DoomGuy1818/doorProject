package handlers

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ServiceHandlers struct {
	service   *service.Service
	validator *validator.Validate
}

func NewServiceHandler(service *service.Service, validator *validator.Validate) *ServiceHandlers {
	return &ServiceHandlers{
		service:   service,
		validator: validator,
	}
}

func (h *ServiceHandlers) CreateService(ctx echo.Context) error {
	serviceDto := new(dto.ServiceDto)

	if err := ctx.Bind(serviceDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(serviceDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	form, err := h.service.CreateService(serviceDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, form)
}
