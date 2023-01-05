package main


/*
Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido 
que al pegarle deberá responder en texto “Hola + nombre + apellido”

El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
La estructura deberá ser como esta:
{
		“nombre”: “Andrea”,
		“apellido”: “Rivas”
}


*/

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
	router := gin.Default()


	// paso los parametros por body -> raw
	router.POST("/saludo", func(ctx *gin.Context) {

		var person Persona

		//Guardamos el json en la variable de tipo person persona: 
		/*Le estoy pasando: 
		{
			"nombre": "aa",
			"apellido": "bb"
		}
		*/
		if err := ctx.BindJSON(&person); err != nil {
			ctx.String(http.StatusInternalServerError, "Bad format")
		} else {
			message := fmt.Sprintf("Hola %s %s", person.Nombre, person.Apellido)
			ctx.String(http.StatusOK, message)
		}
	})

	if err := router.Run(":3030"); err != nil {
		panic(err)
	}

}