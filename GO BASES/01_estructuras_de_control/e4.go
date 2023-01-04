package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Edad de benjamin: %d \n", employees["Benjamin"])

	contador := 0

	for _, valor := range employees {

		if valor > 21 {
			contador++
		}
	}

	fmt.Printf("Cantidad mayores a 21: %d \n", contador)

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println(employees)

}