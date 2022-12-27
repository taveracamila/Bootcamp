package main

import "fmt"


type Person struct{
	ID int
	Name string
	DateOfBirth string
}


type Employee struct{
	Person
	Position string
}

func (e *Employee) PrintEmployee(){
	fmt.Printf("%+v", e)

}



func main(){

	e:=Employee{1,"Luis", "1/1/1998", "gerente"}
	e.PrintEmployee()




	

}