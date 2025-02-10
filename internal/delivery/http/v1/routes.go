package v1

import (
	"doorProject/internal/delivery/http/v1/routes"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	echo    *echo.Echo
	product *routes.ProductRoute
}

func NewRoutes(p *routes.ProductRoute, echo *echo.Echo) *Routes {
	return &Routes{
		echo:    echo,
		product: p,
	}
}

func (r *Routes) InitRoutes() {
	r.product.CreateProduct(r.echo)
}
