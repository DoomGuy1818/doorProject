package server

import (
	v1 "doorProject/internal/delivery/http/v1"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo   *echo.Echo
	routes *v1.Routes
}

func NewServer(routes *v1.Routes, echo *echo.Echo) *Server {
	return &Server{
		echo:   echo,
		routes: routes,
	}
}

func (s Server) Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	s.echo.GET(
		"/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		},
	)

	s.routes.InitRoutes()

	s.echo.Logger.Fatal(s.echo.Start(":" + os.Getenv("PORT")))
}
