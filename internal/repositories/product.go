package repositories

import (
	"fmt"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) GetByID(id string) (string, error) {
	// Simulate a database call
	if id == "" {
		return "", fmt.Errorf("product ID cannot be empty")
	}
	return "Product " + id, nil
}
