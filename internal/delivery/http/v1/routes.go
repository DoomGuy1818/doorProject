package v1

import (
	"doorProject/internal/delivery/http/v1/routes"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	echo           *echo.Echo
	product        *routes.ProductRoute
	color          *routes.ColorRoutes
	category       *routes.CategoryRoute
	worker         *routes.WorkerRoutes
	workerCalendar *routes.WorkerCalendar
}

func NewRoutes(
	p *routes.ProductRoute,
	c *routes.ColorRoutes,
	cat *routes.CategoryRoute,
	w *routes.WorkerRoutes,
	wc *routes.WorkerCalendar,
	echo *echo.Echo,
) *Routes {
	return &Routes{
		echo:           echo,
		product:        p,
		color:          c,
		category:       cat,
		worker:         w,
		workerCalendar: wc,
	}
}

func (r *Routes) InitRoutes() {
	r.product.CreateProduct(r.echo)
	r.color.CreateColor(r.echo)
	r.category.CreateCategory(r.echo)
	r.worker.CreateWorker(r.echo)
	r.workerCalendar.CreateWorkDay(r.echo)
}
