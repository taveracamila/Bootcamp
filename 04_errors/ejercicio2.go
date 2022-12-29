

package main

import (
	"fmt"
	"errors"
)

var MyError = CustomError{"Error: el salario es menor a 10.000"}

func main(){

	salary:=4

	err:=validarSalario(salary)

	if errors.Is(err, MyError) {
		fmt.Printf("Ocurrio un error: %s", err)
	}else{
		fmt.Println("No hubo error")
	}


}



type CustomError struct{
	msg string
}

func (e CustomError) Error() string {
		return e.msg
}


func validarSalario(salary int) (err error){
	if salary<10000{
		err = CustomError{"Error: el salario es menor a 10.000"}
	}
	return 
}


