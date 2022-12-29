package main

import (
	"errors"
	"fmt"
)



func main() {

	var horas int
	fmt.Print("Ingrese horas trabajadas en el mes: ")
	fmt.Scan(&horas)

	var valorHora int
	fmt.Print("Ingrese el valor de la hora trabajada: ")
	fmt.Scan(&valorHora)

	salary, err := calcularSalario(horas, valorHora)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("El salario mensual es de %d\n", salary)



	err2 := validarSalario(salary)
	if err != nil {
		fmt.Println(err2)
	}else{
		fmt.Println("No hubo error al validar el salario")
	}
}



func calcularSalario(horas int, valor int) (salary int, err error) {

	if horas < 0 || horas > 80 {
		err = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}

	salary = horas * valor

	if salary >= 150000 {
		descuento:=float64(salary) * 0.1
		salary -= int(descuento)
	}
	return
}

func validarSalario(salary int) (err error){
	if salary<10000{
		err = fmt.Errorf("Error: el mínimo imponible es de 10000 y el salario ingresado es de: %d", salary)
	}
	return 
}