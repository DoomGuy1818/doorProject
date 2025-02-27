package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type WorkerCalendar struct {
	handler *handlers.WorkerCalendarHandlers
}

func NewWorkerCalendar(handler *handlers.WorkerCalendarHandlers) *WorkerCalendar {
	return &WorkerCalendar{
		handler: handler,
	}
}

func (wc *WorkerCalendar) CreateWorkDay(echo *echo.Echo) {
	echo.POST("/workDay", wc.handler.CreateWorkDay)
}
