package handlers

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/service"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service   *service.ProductService
	validator *validator.Validate
}

func NewProductHandler(serv *service.ProductService, v *validator.Validate) *ProductHandler {
	return &ProductHandler{
		service:   serv,
		validator: v,
	}
}

func (h ProductHandler) CreateProduct(ctx echo.Context) error {

	p := new(models.Product)

	if err := ctx.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(p); err != nil {
		if errors.As(err, &validator.ValidationErrors{}) {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
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
