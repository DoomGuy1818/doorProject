package v1

import (
	"doorProject/internal/delivery/http/v1/routes"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	echo    *echo.Echo
	product *routes.ProductRoute
	color   *routes.ColorRoutes
}

func NewRoutes(p *routes.ProductRoute, c *routes.ColorRoutes, echo *echo.Echo) *Routes {
	return &Routes{
		echo:    echo,
		product: p,
		color:   c,
	}
}

func (r *Routes) InitRoutes() {
	r.product.CreateProduct(r.echo)
	r.color.CreateColor(r.echo)
}
