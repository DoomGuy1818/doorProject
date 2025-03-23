package routes

import (
	"doorProject/internal/config/configs"
	"doorProject/internal/delivery/http/v1/handlers"
	"doorProject/pkg/service"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type WorkerCalendar struct {
	handler     *handlers.WorkerCalendarHandlers
	authService service.AuthService
	middleware  *configs.MiddlewareConfig
}

func NewWorkerCalendar(
	handler *handlers.WorkerCalendarHandlers,
	mdl *configs.MiddlewareConfig,
	auth service.AuthService,
) *WorkerCalendar {
	return &WorkerCalendar{
		handler:     handler,
		authService: auth,
		middleware:  mdl,
	}
}

func (wc *WorkerCalendar) WorkCalendarRoutes(echo *echo.Echo) {
	protectedGroup := echo.Group("/workday")
	protectedGroup.Use(wc.authService.TokenRefresherMiddleware)
	protectedGroup.Use(wc.setupJWTConfig())

	protectedGroup.POST("", wc.handler.CreateWorkDay)
}

func (wc *WorkerCalendar) setupJWTConfig() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey:   wc.middleware.SigningKey,
		TokenLookup:  wc.middleware.TokenLookup,
		ErrorHandler: wc.middleware.ErrorHandler,
	}
	return echojwt.WithConfig(config)
}
