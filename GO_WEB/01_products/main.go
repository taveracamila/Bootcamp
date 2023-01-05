package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"

)


type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}



var Products []Product




func main(){

	/* EJERCICIO 1

		El siguiente paso será crear un archivo main.go donde deberán cargar en una slice, 
		desde un archivo JSON, los datos de productos. 
		Esta slice se debe cargar cada vez que se inicie la API para realizar las distintas consultas.

	*/
	err := cargarJSON("./products.json")
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/ping", Ping)
	router.GET("/products", ListarProductos)
	router.GET("/products/:id", GetProductById)
	router.GET("/products/search", GetProductsPriceGt)

	router.Run(":2020")



	


}



//EJERCICIO 1 
func cargarJSON(path string) (err error){

	obj, err := os.ReadFile(path)
	if err != nil {
		return
	}
	json.Unmarshal(obj, &Products)
	return
}



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

		if prod.Id == id {
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
	for _, item := range Products {
		if priceQuery != 0 && item.Price > float64(priceQuery) {
			response = append(response, item)
		}
	}
	ctx.JSON(200, gin.H{
		"message": "PRODUCTOS ENCONTRADOS:",
		"data":    response,
	})
}


