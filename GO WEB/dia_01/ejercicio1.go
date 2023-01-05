package main 

import "github.com/gin-gonic/gin"


/*
Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
El endpoint deberá ser de método GET
La respuesta de “pong” deberá ser enviada como texto, NO como JSON

*/


func main(){

	router:=gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		/*
		c.JSON(200, gin.H{
			"message":"pong",
		})
		*/
	
		c.String(200, "pong")
	})

	router.Run(":2020")


}