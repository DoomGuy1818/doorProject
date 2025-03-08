package handlers

import (
	"net/http"
	"strconv"
	"time"

	"doorProject/internal/service"

	"github.com/labstack/echo/v4"
)

type AppointmentHandler struct {
	service *service.AppointmentService
}

func NewAppointmentHandler(service *service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		service: service,
	}
}

// GetFreeSlots @Summary Получить свободные слоты
// @Description Возвращает доступные временные интервалы для записи
// @Tags Записи
// @Param worker_id query int true "ID мастера"
// @Param date query string true "Дата (ГГГГ-ММ-ДД)"
// @Success 200 {array} models.TimeSlot
// @Router /slots [get]
func (h *AppointmentHandler) GetFreeSlots(c echo.Context) error {
	// Парсим параметры
	workerID, err := parseUintParam(c, "worker_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("неверный ID мастера"))
	}

	date, err := time.Parse("2006-01-02", c.QueryParam("date"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("неверный формат даты"))
	}

	// Получаем слоты
	slots, err := h.service.GetFreeSlots(workerID, date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, slots)
}

func parseUintParam(c echo.Context, param string) (uint, error) {
	query := c.Param(param)
	masterId, err := strconv.ParseInt(query, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(masterId), nil
}

func errorResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"error": msg,
	}
}
