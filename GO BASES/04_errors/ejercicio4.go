

package main

import (
	"fmt"
	"errors"
)


var MyError = errors.New("Error: el salario es menor a 10.000")

func main(){

	var salary int
	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salary)


	err:=validarSalario(salary)

	if err != nil {
		fmt.Printf("Ocurrio un error: %s \n", err)
	}else{
		fmt.Println("No hubo error")
	}


}




func validarSalario(salary int) (err error){
	if salary<10000{
		err = fmt.Errorf("Error: el mínimo imponible es de 10000 y el salario ingresado es de: %d", salary)
	}
	return 
}
