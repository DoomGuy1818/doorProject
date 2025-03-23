package configs

import "github.com/labstack/echo/v4"

type MiddlewareConfig struct {
	SigningKey   []byte
	TokenLookup  string
	ErrorHandler func(c echo.Context, err error) error
}
