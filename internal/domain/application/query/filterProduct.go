package query

import (
	"log"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
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
	log.Println(minPrice)
	log.Println(maxPrice)
	if maxPrice < minPrice {
		return nil, model2.MinAndMaxPriceCombinationError{}
	}

	return queryEvent.repository.Filter(minPrice, maxPrice)
}
