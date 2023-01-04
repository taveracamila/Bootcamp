package main

import (
	// "encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
	
)

type Cliente struct {
	Legajo int
	Nombre string
	DNI string 
	Telefono   string
	Domicilio string
}

var clientes []*Cliente



func main(){


	cliente := Cliente{
		Legajo: 12,
		Nombre: "Carla",
		DNI: "23237777",
		Telefono: "11898989",
		Domicilio: "muniz 333",
	}

	cliente2 := Cliente{
		Legajo: 88,
		Nombre: "Ramona",
		DNI: "15778987",
		Telefono: "11898989",
		Domicilio: "yapeyu 7776",
	}


	cargarFile()
	agregarCliente(cliente)
	agregarCliente(cliente2)

	


	
}

func cargarFile(){


	defer func() {
		fmt.Println("defer: error de ejecucion")

		if err := recover(); err != nil {
			fmt.Printf("Hubo un error en la ejecución. \nMensaje de error: %v\n", err)
		}

		
	}()


	clientes=nil

	file, err := os.Open("clientes.csv")

	if err != nil {
		panic("Error al abrir el archivo")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		id, _ := strconv.Atoi(items[0])
		
		cliente := Cliente{id, items[1], items[2], items[3], items[4]}
		clientes=append(clientes, &cliente)

		if err != nil {
			fmt.Print("estoy en cargar: ",err)
		}	
	}







			


}










func clienteValidarNoExista(cliente Cliente) (err error) {

	// recupero el panic
	defer func(){
		err:=recover() 

		if err!=nil{

			fmt.Println(err)
		}
	}()

	for _, item := range clientes {

		if item.Legajo == cliente.Legajo {
			
			panic("Error: el cliente ya existe")
		}
	}

	return 

}



func clienteValidarCampos(cliente Cliente)  (err error) {
	if cliente.Legajo == 0 {
		err= errors.New("El legajo no puede ser cero.")
	}
	if cliente.Nombre == "" {
		err=errors.New("Nombre clienteacio")
	}
	if cliente.DNI == "" {
		err= errors.New("DNI no puede ser cero.")
	}
	if cliente.Telefono == "" {
		err= errors.New("Telefono clienteacio")
	}
	if cliente.Domicilio == "" {
		err= errors.New("Domicilio clienteacio.")
	}
	return  err
}



func agregarCliente(cliente Cliente) (err error) {


	err=clienteValidarNoExista(cliente)
	
	if err!=nil{
		fmt.Println(err)
		
	}

	err=clienteValidarCampos(cliente)
	
	if err!=nil {
		fmt.Println("estoy en agregar", err)
	}

	

	text := fmt.Sprintf("\n%d,%s,%s,%s,%s", cliente.Legajo, cliente.Nombre, cliente.DNI, cliente.Telefono, cliente.Domicilio)
	file, err := os.OpenFile("clientes.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)


	if err != nil {
		fmt.Println("estoy en agregar", err)
	}


	if _, err = file.WriteString(text); err != nil {
		fmt.Println("erroooor de escritura", err)
	}else{
		clientes=append(clientes, &cliente)
	}


	// defer
	defer file.Close()

	return 
}


