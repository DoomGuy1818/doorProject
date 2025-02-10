package handlers

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(serv *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: serv,
	}
}

func (h ProductHandler) CreateProduct(ctx echo.Context) error {

	p := new(models.Product)

	if err := ctx.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateProduct(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, p)
}

func UpdateProduct(id int) error {
	return nil
}

func DeleteProduct(id int) error {
	return nil
}

func GetProductById(id int) error {
	return nil
}

func GetProducts() error {
	return nil
}
