package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)






// Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}


//Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
func ListarProductos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Listado de productos", "data": Products})
}


// Crear una ruta /products/	:id que nos devuelva un producto por su id.
func GetProductById(c *gin.Context) {
	// request
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ERROR ", "data": nil})
		return
	}

	// busco porducto
	var response Product
	var flag bool

	for _, prod := range Products {
		if prod.ID == id {
			response = prod
			flag = true
		}
	}



	// response
	if flag {
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

	//var filteredProducts = make([]product.Producto, 0)

	var response []Product
	for _, item := range product.Products {
		if priceQuery != 0 && item.Price > float64(priceQuery) {
			response = append(response, item)
		}
	}
	ctx.JSON(200, gin.H{
		"message": "PRODUCTOS ENCONTRADOS:",
		"data":    response,
	})
}