package query

import (
	"errors"
	"meliarqsoft2/internal/domain"
)

type FilterProductEvent struct {
	repository domain.IProductRepository
}

func NewFilterProductEvent(repository domain.IProductRepository) *FilterProductEvent {
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
