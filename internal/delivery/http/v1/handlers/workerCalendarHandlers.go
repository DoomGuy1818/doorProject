package handlers

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/service"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type WorkerCalendarHandlers struct {
	service   *service.WorkerCalendar
	validator *validator.Validate
}

func NewWorkerCalendarHandlers(s *service.WorkerCalendar, v *validator.Validate) *WorkerCalendarHandlers {
	return &WorkerCalendarHandlers{
		service:   s,
		validator: v,
	}
}

func (wc *WorkerCalendarHandlers) CreateWorkDay(ctx echo.Context) error {
	calendarDto := new(dto.WorkerCalendarDto)

	if err := ctx.Bind(calendarDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := wc.validator.Struct(calendarDto); err != nil {
		if errors.As(err, &validator.ValidationErrors{}) {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
	}

	form, err := wc.service.CreateWorkDay(calendarDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, form)
}
