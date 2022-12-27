package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpuestoSalario(T *testing.T) {

	/*
		//arrarge
		var salario float64 = 40000
		var impuestoEsperado float64 = 0

		//act

		obtenido := impuestoSalario(salario)

		//assert

		assert.Equal(T, impuestoEsperado, obtenido)
	*/

	//arrarge
	var salario float64 = 160000
	var impuestoEsperado float64 = 43200

	//act

	obtenido := impuestoSalario(salario)

	//assert

	assert.Equal(T, impuestoEsperado, obtenido)
}

func TestCalcularPromedio(T *testing.T) {

	//arrarge
	var notas = []float64{10, 6, 5}
	var promedioEsperado float64 = 7

	//act

	obtenido := calcularPromedio(notas...)

	//assert

	assert.Equal(T, promedioEsperado, obtenido)

}

func TestCalcularSalario(T *testing.T) {

	/*
		//arrarge
		minutos := 120
		categoria := "A"
		var valorEsperado float64 = 9000

		//act

		obtenido := calcularSalario(minutos, categoria)

		//assert

		assert.Equal(T, valorEsperado, obtenido)

	*/

	/*

		//arrarge
		minutos := 120
		categoria := "B"
		var valorEsperado float64 = 3600

		//act

		obtenido := calcularSalario(minutos, categoria)

		//assert

		assert.Equal(T, valorEsperado, obtenido)

	*/

	//arrarge
	minutos := 120
	categoria := "C"
	var valorEsperado float64 = 2000

	//act

	obtenido := calcularSalario(minutos, categoria)

	//assert

	assert.Equal(T, valorEsperado, obtenido)
}
