package handlers

import (
	"doorProject/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AppointmentHandler struct {
	service   *service.AppointmentService
	validator *validator.Validate
}

func NewAppointmentHandler(service *service.AppointmentService, validator *validator.Validate) *AppointmentHandler {
	return &AppointmentHandler{
		service:   service,
		validator: validator,
	}
}

func (h *AppointmentHandler) GetFreeSlotsHandler(ctx echo.Context) error {

	workerID, err := strconv.ParseUint(ctx.Param("worker_id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	serviceID, err := strconv.ParseUint(ctx.Param("service_id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	dayStr := ctx.QueryParam("day")
	if dayStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "day parameter is required"})
	}

	day, err := time.Parse("2006-01-02T15:04:05Z07:00", dayStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	slots := h.service.GetAppointmentSlots(day, uint(workerID), uint(serviceID))

	return ctx.JSON(http.StatusOK, slots)
}
