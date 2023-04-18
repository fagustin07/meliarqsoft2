package query

import (
	"errors"
	"meliarqsoft2/internal/domain/model"
)

type FilterProductEvent struct {
	repository model.IProductRepository
}

func NewFilterProductEvent(repository model.IProductRepository) *FilterProductEvent {
	return &FilterProductEvent{
		repository: repository,
	}
}

func (queryEvent FilterProductEvent) Execute(minPrice float32, maxPrice float32) ([]model.Product, error) {
	if maxPrice < minPrice {
		return nil, errors.New("min_price cannot be gt max_price")
	}

	return queryEvent.repository.Filter(minPrice, maxPrice)
}
