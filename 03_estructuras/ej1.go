package main

import (
	"fmt"
)





type Product struct{
	ID int
	Name string
	Price int
	Description string
	Category string
}

func (p *Product) Save(){
	 Products=append(Products, *p)

}

func (p *Product) GetAll(){
	fmt.Printf("%+v", Products)

}

func getById(id int) Product{

	var ret Product

	for index:=range  Products{
		if(Products[index].ID==id){
			ret= Products[index]
			break
		}
	}

	return ret

	

}

var Products=[] Product {Product{3,"Cartuchera", 500, "Dos pisos con cierre", "utiles"}, Product{2,"Botella", 500, "Termica", "Cocina"}}



func main(){

	p:=Product{1,"pizza", 500, "con jamon", "utilcomidaes"}
	p.Save()
	fmt.Printf("%+v", Products)

	prod:=getById(2)
	fmt.Printf(" el producto es   \n %+v", prod)

}