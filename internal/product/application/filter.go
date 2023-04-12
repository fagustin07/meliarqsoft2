package application

import (
	"errors"
	"meliarqsoft2/internal/product/domain"
)

func (service *ProductService) Filter(minPrice float32, maxPrice float32) ([]*domain.Product, error) {
	if maxPrice < minPrice {
		return nil, errors.New("min_price cannot be gt max_price")
	}
	return service.repo.Filter(minPrice, maxPrice)
}
