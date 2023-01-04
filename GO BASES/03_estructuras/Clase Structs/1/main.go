package main

import "fmt"

type Location struct {
	Country  string
	Latitud  float64
	Longitud float64
}
func (l Location) SetLatitud(lat float64) {
	l.Latitud = lat
}
func (l *Location) SetLongitud(long float64) {
	l.Longitud = long
}

type Persona struct {
	FirstName string
	LastName  string
	Age       uint
	Location
}
func (p Persona) Fullname() string {
	return p.FirstName + " " + p.LastName
}

type Medico struct {
	Persona
	Matricula int
}
func (m Medico) FalseSetMatricula(matricula int) {
	m.Matricula = matricula
}
func (m *Medico) SetMatricula(matricula int) {
	m.Matricula = matricula
}

func main() {
	medico := Medico{
		Persona{
			FirstName: "Jane",
			LastName: "Doe",
			Age: 40,
			Location: Location{
				Country: "Argentina",
				Latitud: 50.7,
				Longitud: 80,
			},
		},
		500,
	}

	fmt.Println(medico.Latitud)

	medico.SetLatitud(50)

	fmt.Println(medico.Latitud)

	fmt.Println("nombre", medico.Fullname())

	medico.Persona.FirstName = "---"

	fmt.Println("nombre", medico.Fullname())

}

