package product

import (
	"errors"
	"02_products_api/internal/domain"
)

var (
	ErrAlreadyExist = errors.New("already exist")
)




func CreateProduct(p domain.Product) error {


	if err := ExisteCodeValue(p.CodeValue); err != nil {
		return err
	}


	AgregarProduct(p)
	return nil

	
}

func GuardarJSON(productos [] domain.Product){
	for _, item := range productos {
		AgregarProduct(item)
	}
}


func GetProducts() [] domain.Product{
	return GetAll()
}

func GetProductById(id int) (domain.Product, error){
	return GetProduct(id)
}



func GetProductsPriceGt(price int) [] domain.Product{


	return GetProductsByPriceGt(price)


}


