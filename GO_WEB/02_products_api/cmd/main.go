package main

import (
	"02_products_api/cmd/handlers"
	"02_products_api/internal/product"
	"02_products_api/internal/domain"

	"encoding/json"
	"os"
	"github.com/gin-gonic/gin"

)

func main(){

	err := cargarJSON("./products.json")
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	/*
	router.GET("/ping", Ping)
	router.GET("/products", ListarProductos)
	router.GET("/products/:id", GetProductById)
	router.GET("/products/search", GetProductsPriceGt)

	*/
	router.POST("/products", handlers.AgregarProduct)
	router.GET("/products", handlers.GetAll)
	router.GET("/products/:id", handlers.GetProductById)
	router.GET("/products/search", handlers.GetProductsPriceGt)





	router.Run(":2020")
}



func cargarJSON(path string) (err error){

	obj, err := os.ReadFile(path)
	if err != nil {
		return
	}

	var aux [] domain.Product

	json.Unmarshal(obj, &aux)


	product.GuardarJSON(aux)

	return
}