package handlers

import (
	"context"
	"doorProject/internal/delivery/http/v1/handlers/dto"
	"doorProject/internal/repository"
	"doorProject/pkg/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlers struct {
	auth       service.AuthService
	workerRepo repository.WorkerRepositoryInterface
	publisher  service.MessagePublisherInterface
}

func NewAuthHandlers(
	auth service.AuthService,
	w repository.WorkerRepositoryInterface,
	publisher service.MessagePublisherInterface,
) *AuthHandlers {
	return &AuthHandlers{auth: auth, workerRepo: w, publisher: publisher}
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

func (h *AuthHandlers) Register(c echo.Context) error {
	registerDto := new(dto.RegisterDto)

	if err := c.Bind(registerDto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if registerDto.Password != registerDto.RepeatPassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Пароли не совпадают")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(registerDto.Password),
		bcrypt.DefaultCost,
	)

	fmt.Println(string(hashedPassword))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h.publisher.PublishMessage(ctx, registerDto.Login, "sendEmail")

	return echo.NewHTTPError(http.StatusCreated, "successfully registered. Check your email")
}
