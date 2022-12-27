package main

import (
	"fmt"
)

func main() {

	//ej1
	var salario float64
	fmt.Println("Ingrese su salario")

	fmt.Scan(&salario)
	impuesto := impuestoSalario(salario)

	fmt.Printf("El impuesto es: %f \n", impuesto)

	//ej2

	promedio := calcularPromedio(6, 10)

	fmt.Printf("El promedio es: %f \n", promedio)

}

func impuestoSalario(salario float64) float64 {
	if salario > 50000 && salario < 150000 {
		return salario * 0.17
	} else if salario > 150000 {
		return salario * 0.27

	} else {
		return 0
	}
}

// ej2
func calcularPromedio(calificaciones ...float64) float64 {
	var suma float64 = 0
	for index := range calificaciones {
		suma += calificaciones[index]
	}

	return suma / float64(len(calificaciones))

}

// ej3
func calcularSalario(minutosTrabajados int, categoria string) (salario float64) {
	horas := float64(minutosTrabajados) / float64(60)

	switch categoria {
	case "C":
		salario = horas * 1000
	case "B":
		aux := 1500 * horas
		salario = aux + aux*0.2
	case "A":
		aux := 3000 * horas
		salario = aux + aux*0.5
	}

	return

}

/*

func operation(op string) (func(operandos ...int) float64, error) {
	switch op {
	case "minimum":
		return min, nil
	case "maximum":
		return max, nil
	case "average":
		return average, nil
	}

	return nil, errors.New("No such operator")
}

func min(operandos ...int) float64 {
	min := operandos[0]
	for _, value := range operandos {
		if value < min {
			min = value
		}
	}

	return float64(min)
}

func max(operandos ...int) float64 {
	max := operandos[0]
	for _, value := range operandos {

	}
}

func salario(minutos int, categoria string) float64 {
	horas := float64(minutos) / float64(60)
	switch categoria {
	case "A":
		return 3000 * horas * 1.50
	case "B":
		return 1500 * horas * 1.20
	case "C":
		return 1000 * horas
	}
	return 0
}



var foodAmount = map[string]int{
	"dog":       10000,
	"cat":       5000,
	"hamster":   250,
	"tarantula": 150,
}

func animal(animal string) (func(amount int) int, error) {
	if _, ok := foodAmount[animal]; ok {
		return func(amount int) int { return amount * foodAmount[animal] }, nil
	}
	return nil, errors.New("Invalid animal")
}







func impuesto(salario int64) float64 {
	impuesto := 0.0
	if salario > 50000 {
		impuesto += 0.17
	}
	if salario > 150000 {
		impuesto += 0.10
	}

	return float64(salario) * impuesto
}

func promedio(notas ...uint) float64 {
	var sum uint = 0
	for _, nota := range notas {
		sum += nota
	}

	return float64(sum) / float64(len(notas))
}
*/
