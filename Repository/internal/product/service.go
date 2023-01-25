package product

import (
	"context"
	// "database/sql"
	"errors"
	 // "time"

	"Repository/internal/domain"

	// "github.com/go-playground/validator/v10"
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