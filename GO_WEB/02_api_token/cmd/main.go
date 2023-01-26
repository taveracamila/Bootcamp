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


	var products []domain.Product
	err := cargarJSON("./products.json", &products)
	if err != nil {
		panic(err)
	}

	repo := product.NewRepository(products)
	service := product.NewService(repo)
	ph := handlers.NewProductHandler(*service)

	router := gin.Default()
	/*
	router.GET("/ping", Ping)
	router.GET("/products", ListarProductos)
	router.GET("/products/:id", GetProductById)
	router.GET("/products/search", GetProductsPriceGt)

	*/

	pr := router.Group("/products")
	{
		pr.POST("", ph.Create())
		pr.GET("", ph.GetAll())
		pr.GET("/:id", ph.GetProductById())
		pr.GET("/search", ph.GetProductsPriceGt())
		pr.PUT("/:id", ph.Update())
		pr.PATCH("/:id", ph.UpdatePrice())
		pr.DELETE("/:id", ph.Delete())


	
	}

	/*
	router.POST("/products", ph.Create())
	router.GET("/products", ph.GetAll())
	router.GET("/products/:id", ph.GetProductById())
	router.GET("/products/search", ph.GetProductsPriceGt())
	products.PUT("/:id", ph.Update())

	*/






	router.Run(":2020")
}



func cargarJSON(path string, list *[]domain.Product) (err error){

	obj, err := os.ReadFile(path)
	if err != nil {
		return
	}
	json.Unmarshal(obj, &list)
 	return
}