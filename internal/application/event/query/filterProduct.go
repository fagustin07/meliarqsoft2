package query

import (
	"errors"
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type FilterProductEvent struct {
	repository ports.IProductRepository
}

func NewFilterProductEvent(repository ports.IProductRepository) *FilterProductEvent {
	return &FilterProductEvent{
		repository: repository,
	}
}

func (queryEvent FilterProductEvent) Execute(minPrice float32, maxPrice float32) ([]*domain.Product, error) {
	if maxPrice < minPrice {
		return nil, errors.New("min_price cannot be gt max_price")
	}

	return queryEvent.repository.Filter(minPrice, maxPrice)
}
