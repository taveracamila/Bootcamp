package handlers

import (
	"net/http"

	"02_products_api/internal/domain"
	"02_products_api/internal/product"

	"github.com/gin-gonic/gin"
	"time"
	"errors"
	"strconv"

)


// recibe peticiones del front 
// valida params



func AgregarProduct(c *gin.Context) {
	var req domain.Product 

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	

	if err := verificarFecha(req.Expiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al verificar la fecha ",
		})
		return
	}

	if err := verificarVacios(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error verificando vacios",
		})
		return
	}


	product.CreateProduct(req)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
		"data":    req,
	})
}




func verificarFecha(date string) error {
	format := "02/01/2006"
	_, err := time.Parse(format, date)
	return err
}

func verificarVacios(prod domain.Product) error {
	if prod.Price == 0 {
		return errors.New("Price vacio")}
	if prod.Name == "" {
		return errors.New("Name vacio")
	}
	if prod.Expiration == "" {
		return errors.New("Expiration vacio")
	}
	if prod.CodeValue == "" {
		return errors.New("CodeValue  vacio")
	}
	if prod.Quantity == 0 {
		return errors.New("Quantity vacio")
	}
	return nil
}


//Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
func GetAll(c *gin.Context) {
	data:=product.GetAll()
	c.JSON(http.StatusOK, gin.H{"message": "Listado de productos", "data": data})
}


func GetProductById(c *gin.Context) {
	// request
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ERROR ", "data": nil})
		return
	}




	response, err_func:=product.GetProductById(id)
	if err_func == nil {
		c.JSON(http.StatusOK, gin.H{"message": "PRODUCTO ENCONTRADO", "data": response})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "NO SE ENCONTRO EL PRODUCTO", "data": nil})
	}
}



// Crear una ruta /products/search que nos permita buscar por parámetro los productos 
// cuyo precio sean mayor a un valor priceGt.

func GetProductsPriceGt(ctx *gin.Context) {
	priceQuery, err := strconv.Atoi(ctx.Query("price"))

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "NO SE ENCONTRARON PRODUCTOS",
			"data":    nil,
		})
		return
	}

	response:=product.GetProductsPriceGt(priceQuery)

	ctx.JSON(200, gin.H{
		"message": "PRODUCTOS ENCONTRADOS:",
		"data":    response,
	})
}













