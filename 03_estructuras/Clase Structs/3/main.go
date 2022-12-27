package main

import (
	"errors"
	"fmt"
)

// models
type Item struct {
	Title, Description string
}
type ItemStore struct {
	Item
	Price    float64
	Quantity float64
}

// controller
type Shop struct {
	// config
	transactions int
	limit		 int
	// services
	storage []ItemStore
}
func NewShop(limit int, storage []ItemStore) *Shop {
	return &Shop{
		transactions: 1,
		limit: limit,
		storage: storage,
	}
}

func (s *Shop) Available() bool {
	return s.transactions < s.limit
}
func (s *Shop) IsEmpty() bool {
	return len(s.storage) == 0
}
func (s *Shop) CargaStock(items ...ItemStore) error {
	if !s.Available() {
		return errors.New("not available")
	}

	s.storage = append(s.storage, items...)

	s.transactions++

	return nil
}

func main() {
	storage := []ItemStore{}
	
	shop := NewShop(3, storage)
	
	err := shop.CargaStock(ItemStore{Item{"manzana", ""}, 500, 1})
	if err != nil {
		panic(err)
	}
	err = shop.CargaStock(ItemStore{Item{"manzana", ""}, 500, 1})
	if err != nil {
		panic(err)
	}

	fmt.Println(shop.storage)
}