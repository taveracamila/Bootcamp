
package product

import (
	"errors"
	"02_products_api/internal/domain"

)


type Repository interface {
	GetAll() []domain.Product
	GetProductById(id int) (p domain.Product, err error )
	GetProductsPriceGt(precio int ) [] domain.Product
	Create(p domain.Product)
	ExisteCodeValue(code string) error
	Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (p domain.Product, err error)
}
	

type repository struct {

	products []domain.Product
}



var LastId int

func NewRepository(listProducts []domain.Product) *repository {
	return &repository{listProducts}
}


func (r *repository) ExisteCodeValue(code string) error {
	for _, item := range r.products {
		if item.CodeValue == code {
			return errors.New("CodeValue existente")
		}
	}
	return nil
}


func (r *repository) Create(p domain.Product) {


	r.products = append(r.products, p)

}


func (r *repository) GetAll()  [] domain.Product{

	return r.products

}


func (r *repository) GetProductById(id int) (p domain.Product, err error ){

	for _, prod := range r.products {

		if prod.Id == id {
			p=prod
			return
		}
	}


	err=errors.New("No existe el producto con ese Id")

	return 


}





func (r *repository) GetProductsPriceGt(precio int ) [] domain.Product{

	var response [] domain.Product
	for _, item := range r.products {
		if precio != 0 && item.Price > float64(precio) {
			response = append(response, item)
		}
	}

	return response

}





func (r *repository) Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (p domain.Product, err error) {
	
	//p := domain.Product{Id:id, Name: name, Quantity: quantity, CodeValue: codeValue, IsPublished: isPublished, Expiration: expiration, Price: price}
	p.Id=id
	p.Name=name
	p.Quantity=quantity
	p.CodeValue=codeValue
	p.IsPublished=isPublished
	p.Expiration=expiration
	p.Price=price

	var flag bool



	for _, prod := range r.products {

		if prod.Id == p.Id {
			prod=p
			flag=true
			break
		}
	}

	
	if !flag {
		err = errors.New("No se encontro producto con el id mencionado")
	}

	return p, err
}


/*
// updatePrice actualiza el precio de un producto
func (r *repository) UpdatePrice(id int, price float64) (domain.Product, error) {
	var p domain.Product
	updated := false
	for i := range r.list {
		if r.list[i].Id == id {
			r.list[i].Price = price
			updated = true
			p = r.list[i]
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("couldn't find a product with the id: %d", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range r.list {
		if r.list[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("couldn't find a product with the id: %d", id)
	}
	r.list = append(r.list[:index], r.list[index+1:]...)
	return nil
}

*/





