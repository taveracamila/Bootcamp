package main

import "fmt"

func main() {
	var nmes int
		fmt.Println("Ingrese un numero de mes")
		fmt.Scan(&nmes)

		switch nmes {

		case 1:
			fmt.Println("Enero")
		case 2:
			fmt.Println("Febero")
		case 3:
			fmt.Println("Marzo")
		case 4:
			fmt.Println("Abril")
		case 5:
			fmt.Println("Mayo")
		case 6:
			fmt.Println("Junio")
		case 7:
			fmt.Println("Julio")
		case 8:
			fmt.Println("Agosto")
		case 9:
			fmt.Println("Septiembre")
		case 10:
			fmt.Println("Octubre")
		case 11:
			fmt.Println("Noviembre")
		case 12:
			fmt.Println("Diciembre")
		default:
			fmt.Println("mes invalido")

		}

}