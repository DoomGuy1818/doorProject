package handlers

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ServiceHandlers struct {
	service *service.Service
}

func NewServiceHandler(service *service.Service) *ServiceHandlers {
	return &ServiceHandlers{
		service: service,
	}
}

func (h *ServiceHandlers) CreateService(ctx echo.Context) error {
	serviceDto := new(dto.ServiceDto)

	if err := ctx.Bind(serviceDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(serviceDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	form, err := h.service.CreateService(serviceDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, form)
}
