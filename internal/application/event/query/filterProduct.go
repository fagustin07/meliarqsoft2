package query

import (
	"errors"
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type FilterProduct struct {
	repository ports.IProductRepository
}

func (queryEvent FilterProduct) Execute(minPrice float32, maxPrice float32) ([]*domain.Product, error) {
	if maxPrice < minPrice {
		return nil, errors.New("min_price cannot be gt max_price")
	}

	return queryEvent.repository.Filter(minPrice, maxPrice)
}
