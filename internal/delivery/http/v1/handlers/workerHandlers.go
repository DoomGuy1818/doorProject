package handlers

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/service"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type WorkerHandlers struct {
	workerService service.WorkerService
	validator     validator.Validate
}

func NewWorkerHandlers(workerService service.WorkerService, validator validator.Validate) *WorkerHandlers {
	return &WorkerHandlers{
		workerService: workerService,
		validator:     validator,
	}
}

func (w WorkerHandlers) CreateWorker(ctx echo.Context) error {

	workerDto := new(dto.WorkerDTO)

	if err := ctx.Bind(workerDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := w.validator.Struct(workerDto); err != nil {
		if errors.As(err, &validator.ValidationErrors{}) {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
	}

	form, err := w.workerService.CreateWorker(workerDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, form)
}
