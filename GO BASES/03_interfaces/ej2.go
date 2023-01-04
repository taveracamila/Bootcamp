package main

import "fmt"

const(
	pequeno="pequeno"
	mediano="mediano"
	grande="grande"
)

type Producto interface{
	precio() float64
}

func factory(typeProducto string, precio float64) Producto {

	
	switch(typeProducto){
	case pequeno:
		return ProductoPequeno{costo:precio}
	case mediano:
		return ProductoMediano{costo:precio}
	case grande:
		return ProductoGrande{costo:precio}

	}
	return nil
}


type ProductoPequeno struct{
	costo float64
}

type ProductoMediano struct{
	costo float64
	
}

type ProductoGrande struct{
	costo float64
	
}


func (p ProductoPequeno) precio() float64{
	return  p.costo

}

func (p ProductoMediano) precio() float64{
	return  p.costo + p.costo*0.03

}

func (p ProductoGrande) precio() float64{
	return  p.costo + p.costo*0.06 + 2500

}


func main(){

	i:=factory("pequeno", 1000)
	fmt.Println(i.precio())

}