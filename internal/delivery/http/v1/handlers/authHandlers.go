package handlers

import (
	"context"
	"doorProject/internal/delivery/http/v1/handlers/dto"
	service2 "doorProject/internal/service"
	"doorProject/pkg/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlers struct {
	auth          service.AuthService
	workerService service2.WorkerService
	publisher     service.MessagePublisherInterface
}

func NewAuthHandlers(
	auth service.AuthService,
	w service2.WorkerService,
	publisher service.MessagePublisherInterface,
) *AuthHandlers {
	return &AuthHandlers{auth: auth, workerService: w, publisher: publisher}
}

func (h *AuthHandlers) SignIn(c echo.Context) error {
	credDto := new(dto.SignInDto)

	if err := c.Bind(credDto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	worker, err := h.workerService.GetUser(credDto)
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

func (h *AuthHandlers) VerifyAccount(c echo.Context) error {
	email := c.QueryParam("email")

	status, err := strconv.ParseBool(c.QueryParam("status"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	worker, err := h.workerService.UpdateWorkerStatus(email, status)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"worker": worker})
}

func (h *AuthHandlers) Refresh(c echo.Context) error {
	access, err := h.auth.RefreshAccessToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"access-token": access})
}
