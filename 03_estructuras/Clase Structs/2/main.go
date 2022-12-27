package main

import "fmt"

type Pizza struct {
	Harina float64
	Salsa  float64
	Queso  float64
}

func (p Pizza) Cost() float64 {
	return p.Harina + p.Queso + p.Salsa
}

type PizzaWithAnana struct {
	Pizza
	Anana float64
}

func (p PizzaWithAnana) Cost() float64 {
	return p.Anana + p.Pizza.Cost()
}

func main() {
	pizza := PizzaWithAnana{
		Pizza: Pizza{5, 6, 7},
		Anana: 8,
	}

	fmt.Println(pizza.Cost())
}