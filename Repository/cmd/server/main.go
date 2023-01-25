package main

import (
	"database/sql"
	"log"
	"Repository/cmd/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	databaseConfig := &mysql.Config{
		User:      "root",
		Passwd:    "",
		Addr:      "localhost:3306",
		DBName:    "my_db",
		ParseTime: true,
	}

	db, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}


	if err = db.Ping(); err != nil {
		log.Println("ERROR EN EL PING ")
		panic(err)
	}

	log.Println("!!!!!!!!! SE ESTABLECIO LA CONEXION !!!!!!!!!!!!! ")

	eng := gin.Default()
	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}