package handlers

import (
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/repository"
	"doorProject/pkg/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlers struct {
	auth       service.AuthService
	workerRepo repository.WorkerRepositoryInterface
}

func NewAuthHandlers(auth service.AuthService, w repository.WorkerRepositoryInterface) *AuthHandlers {
	return &AuthHandlers{auth: auth, workerRepo: w}
}

func (h *AuthHandlers) SignIn(c echo.Context) error {
	credDto := new(dto.SignInDto)

	if err := c.Bind(credDto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	worker, err := h.workerRepo.FindUserByUsername(credDto.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(worker.Password), []byte(credDto.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := h.auth.GenerateTokenAndSetCookie(worker, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
