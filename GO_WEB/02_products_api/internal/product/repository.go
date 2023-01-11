
package product

import (
	"errors"
	"02_products_api/internal/domain"

)



	var Products [] domain.Product
var LastId int


func ExisteCodeValue(code string) error {
	for _, item := range Products {
		if item.CodeValue == code {
			return errors.New("CodeValue existente")
		}
	}
	return nil
}


func AgregarProduct(p domain.Product) {


	LastId++
	p.Id = LastId
	Products = append(Products, p)

}


func GetAll()  [] domain.Product{

	return Products

}


func GetProduct(id int) (p domain.Product, err error ){

	for _, prod := range Products {

		if prod.Id == id {
			p=prod
			return
		}
	}


	err=errors.New("No existe el producto con ese Id")

	return 


}





func GetProductsByPriceGt(precio int ) [] domain.Product{

	var response [] domain.Product
	for _, item := range Products {
		if precio != 0 && item.Price > float64(precio) {
			response = append(response, item)
		}
	}

	return response

}





