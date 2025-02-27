package routes

import (
	"doorProject/internal/delivery/http/v1/handlers"

	"github.com/labstack/echo/v4"
)

type WorkerRoutes struct {
	handler handlers.WorkerHandlers
}

func NewWorkerRoutes(h handlers.WorkerHandlers) *WorkerRoutes {
	return &WorkerRoutes{
		handler: h,
	}
}

func (w *WorkerRoutes) CreateWorker(echo *echo.Echo) {
	echo.POST("/workers", w.handler.CreateWorker)
}
