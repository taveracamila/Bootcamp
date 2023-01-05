
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taveracamila/GO WEB/01_products/handlers"
	"github.com/renzobalbo/goWeb/dia1/E2.2/handlers"
)




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
	router.GET("/ping", handlers.Ping)
	router.GET("/products", handlers.ListarProductos)
	router.GET("/products/:id", handlers.GetProductById)
	router.GET("/products/search", handlers.GetProductsPriceGt)




	


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


