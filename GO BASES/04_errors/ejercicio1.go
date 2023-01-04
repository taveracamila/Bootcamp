package main

import (
	"fmt"
)


func main(){

	salary:=10999922

	if salary<150000{

		e:=&CustomError{"Error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(e)

	}else{
		fmt.Println("Debe pagar impuesto")

	}



}


type CustomError struct{
	msg string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
 
}



