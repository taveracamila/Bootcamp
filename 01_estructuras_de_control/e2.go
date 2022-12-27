package main

import "fmt"

func main() {
	var edad int
	var antiguedad int
	var empleado bool
	var sueldo float64

	fmt.Println("Ingrese la edad")
	fmt.Scan(&edad)

	if edad >= 22 {

		fmt.Println("Es empleado? Ingrese true o false")
		fmt.Scan(&empleado)

		if empleado {

			fmt.Println("Ingrese la antiguedad")
			fmt.Scan(&antiguedad)

			if antiguedad > 1 {

				fmt.Println("Ingrese el sueldo")
				fmt.Scan(&sueldo)

				if sueldo >= 100000 {
					fmt.Println("se otorga prestamo sin interes")

				} else {
					fmt.Println("se otorga prestamo con interes")

				}

			} else {
				fmt.Println("Debe tener ams de un año de antiguedad")

			}

		}

	} else {
		fmt.Println("Tiene que ser mayor de 22")

	}

}
