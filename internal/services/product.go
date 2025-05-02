package services

import "github.com/tranduckhuy/go-ecommerce-backend-api/internal/repositories"

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) GetByID(id string) (string, error) {
	product, err := p.productRepository.GetByID(id)
	if err != nil {
		return "", err
	}
	return product, nil
}
