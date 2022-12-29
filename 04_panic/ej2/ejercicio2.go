package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Ejecución finalizada.")
		}
	}()

	file, err := os.Open("customers.txt")
	
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado.")
	}

	defer file.Close()

	fmt.Println(file)
	content, err := ioutil.ReadAll(file)
	fmt.Print(string(content))
	fmt.Println()

}