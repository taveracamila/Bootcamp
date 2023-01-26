package product

import (
	"errors"
	"02_products_api/internal/domain"
)

var (
	ErrAlreadyExist = errors.New("already exist")
)

type IService interface {
	GetAll() []domain.Product
	GetProductById(id int) (p domain.Product, err error )
	GetProductsPriceGt(precio int ) [] domain.Product
	Create(p domain.Product)
	
	Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error)
	UpdatePrice(id int, price float64) (domain.Product, error)
	Delete(id int) error


}

type Service struct {
	r Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}



func (s *Service) Create(p domain.Product) error {


	if err := s.r.ExisteCodeValue(p.CodeValue); err != nil {
		return err
	}

	



	s.r.Create(p)
	return nil

	
}




func (s *Service) GetAll() [] domain.Product{
	return s.r.GetAll()
}

func (s *Service) GetProductById(id int) (domain.Product, error){
	return s.r.GetProductById(id)
}



func (s *Service) GetProductsPriceGt(price int) [] domain.Product{


	return s.r.GetProductsPriceGt(price)


}


func (s *Service) Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error) {
	return s.r.Update(id, name, quantity, codeValue, isPublished, expiration, price)
}


func (s *Service) UpdatePrice(id int, price float64) (domain.Product, error) {
	return s.r.UpdatePrice(id, price)
}

func (s *Service) Delete(id int) error {
	return s.r.Delete(id)
}

