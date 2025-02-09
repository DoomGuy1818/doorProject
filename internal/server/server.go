package server

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()
	e.GET(
		"/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		},
	)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
