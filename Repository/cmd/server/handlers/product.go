package handlers

import (
	"errors"
	"net/http"
	"strconv"

	// "Repository/internal/domain"
	"Repository/internal/product"
	"Repository/pkg/web"

	"github.com/gin-gonic/gin"
)

// Errors
var (
	ErrProductInternalServer = errors.New("internal server error")
)

type Product struct {
	productService product.Service
}

func NewProduct(service product.Service) *Product {
	return &Product{service}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.productService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, product.ErrInternalServer.Error())
			return
		}
		web.Success(c, http.StatusOK, products)
	}
}


func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "El Id ingresado no es valido")
			return
		}

		response, err := p.productService.GetOne(c, id)

		switch err {
		case product.ErrNotFound:
			web.Error(c, 404, product.ErrInvalidProductData.Error())
			return
		case nil:
			web.Success(c, 200, response)
			return
		default:
			web.Error(c, http.StatusInternalServerError, err.Error())

		}

	}
}