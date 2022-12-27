/*
	TYPE ASSERTION 

	var i interface{} = "hello"
	s, ok:= i.(string)
	fmt.Println(s, ok)

	*/



package main

import "fmt"


type Alumnos struct{
	Nombre string
	Apellido string
	DNI string
	Fecha string
}



func (a *Alumnos) detalle(){
	fmt.Printf("Nombre: [ %s ] \n Apellido: [ %s ] \n DNI: [ %s ] \n Fecha: [ %s ] \n", a.Nombre, a.Apellido, a.DNI, a.Fecha)

}



func main(){

	a:=Alumnos{"Camila", "Tavera", "123123", "4.4.2000"}
	a.detalle()



	

}