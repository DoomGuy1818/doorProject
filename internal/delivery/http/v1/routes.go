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
	client         *routes.ClientRoutes
	service        *routes.ServiceRoutes
	cart           *routes.CartRoutes
	appointment    *routes.AppointmentRoutes
	auth           *routes.AuthRoutes
}

func NewRoutes(
	p *routes.ProductRoute,
	c *routes.ColorRoutes,
	cat *routes.CategoryRoute,
	w *routes.WorkerRoutes,
	wc *routes.WorkerCalendar,
	cl *routes.ClientRoutes,
	s *routes.ServiceRoutes,
	crt *routes.CartRoutes,
	a *routes.AppointmentRoutes,
	auth *routes.AuthRoutes,
	echo *echo.Echo,
) *Routes {
	return &Routes{
		echo:           echo,
		product:        p,
		color:          c,
		category:       cat,
		worker:         w,
		workerCalendar: wc,
		client:         cl,
		service:        s,
		cart:           crt,
		appointment:    a,
		auth:           auth,
	}
}

func (r *Routes) InitRoutes() {
	r.product.CreateProduct(r.echo)
	r.color.CreateColor(r.echo)
	r.category.CreateCategory(r.echo)
	r.worker.CreateWorker(r.echo)
	r.workerCalendar.WorkCalendarRoutes(r.echo)
	r.client.CreateClient(r.echo)
	r.service.CreateService(r.echo)
	r.cart.CreateCart(r.echo)
	r.appointment.GetFreeSlots(r.echo)
	r.appointment.CreateAppointments(r.echo)
	r.auth.SighIn(r.echo)
}
