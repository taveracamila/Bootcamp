package handlers

import (
	"errors"
	"net/http"
	"strconv"

	 "Repository/internal/domain"
	"Repository/internal/product"
	"Repository/pkg/web"

	"github.com/gin-gonic/gin"

	"fmt"
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




func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var prod domain.Product
		// check json type
		if err := c.ShouldBindJSON(&prod); err != nil {
			fmt.Println("rompio en el should bind json")

			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		pCreated, err := p.productService.Store(c, prod.Name, prod.Quantity, prod.CodeValue, prod.IsPublished, prod.Expiration, prod.Price)

		if err== product.ErrInvalidProductData {
			web.Error(c, http.StatusInternalServerError, "invalid product data")

		}else if err== product.ErrProductCodeAlreadyExists {
			web.Error(c, http.StatusInternalServerError, "ya existe")

		}

		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())

		}

		web.Success(c, http.StatusCreated, pCreated)
	}
}

/*
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		var prod domain.Product
		// check json type
		if err := c.ShouldBindJSON(&prod); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		prod, err = p.service.Update(c, prod, id)
		if err != nil {
			if errors.Is(err, product.ErrProductRegistered) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			} else if errors.Is(err, product.ErrNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, ErrProductInternalServer.Error())
			return
		}
		web.Success(c, http.StatusOK, prod)
	}
}


*/

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		err = p.productService.Delete(c, id)
		if err != nil {
			web.Error(c, http.StatusInternalServerError,err.Error())

		}
		web.Success(c, http.StatusNoContent, p)
	}
}