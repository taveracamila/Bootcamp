package domain

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	// Expiration  time.Time `json:"expiration" time_format:"2006-01-02" time_utc:"1"`
	Price       float64   `json:"price"`
}