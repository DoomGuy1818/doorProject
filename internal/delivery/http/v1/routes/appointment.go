package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type AppointmentRoutes struct {
	Handler *handlers.AppointmentHandler
}

func NewAppointmentRoutes(handler *handlers.AppointmentHandler) *AppointmentRoutes {
	return &AppointmentRoutes{Handler: handler}
}

func (r *AppointmentRoutes) GetFreeSlots(e *echo.Echo) {
	e.GET("/worker_id/:id/slots", r.Handler.GetFreeSlots)
}
