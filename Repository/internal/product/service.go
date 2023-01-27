package product

import (
	"context"
	// "database/sql"
	"errors"
	 "time"

	"Repository/internal/domain"

	 "github.com/go-playground/validator/v10"
	 "fmt"
)


var (
	ErrNotFound                 = errors.New("product not found")
	ErrInvalidProductData       = errors.New("Invalid product data")
	ErrProductCodeAlreadyExists = errors.New("Product code already exists")
	ErrInternalServer           = errors.New("We have some internal troubles, try again later")
	ErrProductCodeConflict      = errors.New("The new product code already exists")
)


type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetOne(ctx context.Context, id int) (domain.Product, error)
	
	Store(ctx context.Context, name string, quantity int, codeValue string, isPublished bool, expiration time.Time, price float64) (domain.Product, error)	
	Delete(ctx context.Context, id int) error 

}

type service struct {
	r Repository
}

func NewService(repo *Repository) Service {
	return &service{*repo}
}



func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if products == nil {
		return []domain.Product{}, nil
	}
	return products, nil
}


func (s *service) GetOne(ctx context.Context, id int) (domain.Product, error) {

	obj, err := s.r.GetOne(ctx, id)

	if (obj == domain.Product{}) {
		return obj, ErrNotFound
	}

	if err != nil {
		return obj, err
	}

	return obj, nil
}




func (s *service) Store(ctx context.Context, name string, quantity int, codeValue string, isPublished bool, expiration time.Time, price float64) (domain.Product, error) {
	
	fmt.Println("rompio en el service . store ")


	obj := domain.Product{
		Name:         name,
		Quantity: quantity,
		CodeValue:	codeValue,
		IsPublished:     isPublished,
		Expiration:   expiration,
		Price:	price,
	}

	if err := validator.New().Struct(&obj); err != nil {
		return domain.Product{}, ErrInvalidProductData
	}

	if(s.r.Exists(ctx, obj.CodeValue)){
		return domain.Product{}, ErrProductCodeAlreadyExists
	}

	return s.r.Store(ctx, obj)


}




/*

func (s *service) Update(ctx context.Context, id int, cid *int, companyName *string, address *string, telephone *string) (domain.Seller, error) {

	aux, err := s.r.Get(ctx, id)
	if err != nil {
		return domain.Seller{}, ErrNotFound
	}

	//chequear esto
	if cid != nil {

		if aux.CID == *cid || s.r.Exists(ctx, *cid) == false {
			aux.CID = *cid
		} else {
			return domain.Seller{}, ErrCidExists
			fmt.Println("estot en err cid exists")

		}

	}

	if companyName != nil {
		aux.CompanyName = *companyName

	}

	if address != nil {
		aux.Address = *address

	}

	if telephone != nil {
		aux.Telephone = *telephone

	}
	

	if err := validator.New().Struct(&aux); err != nil {

		return domain.Seller{}, ErrInvalidSellerData
	}

	// exist := s.r.Exists(ctx, *cid)

	if err := s.r.Update(ctx, aux); err != nil {

		return domain.Seller{}, ErrDb
	}

	return aux, nil

}

*/


func (s *service) Delete(ctx context.Context, id int) error {

	return s.r.Delete(ctx, id)

}
