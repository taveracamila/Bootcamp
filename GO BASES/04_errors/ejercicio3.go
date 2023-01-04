

package main

import (
	"fmt"
	"errors"
)


var MyError = errors.New("Error: el salario es menor a 10.000")

func main(){

	salary:=4

	err:=validarSalario(salary)

	if errors.Is(err, MyError) {
		fmt.Printf("Ocurrio un error: %s", err)
	}else{
		fmt.Println("No hubo error")
	}


}




func validarSalario(salary int) (err error){
	if salary<10000{
		err = MyError
	}
	return 
}
