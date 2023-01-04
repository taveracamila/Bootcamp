package main

import (
	"fmt"
	"os"
)


func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Ejecución finalizada.")
		}
	}()

	_, err := os.Open("customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado.")
	}
}