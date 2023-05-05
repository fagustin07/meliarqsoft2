package query

import (
	"meliarqsoft2/internal/domain/model"
)

type FindProductEvent struct {
	repository model.IProductRepository
}

func NewFindProductEvent(repository model.IProductRepository) *FindProductEvent {
	return &FindProductEvent{
		repository: repository,
	}
}

func (queryEvent FindProductEvent) Execute(name string, category string) ([]model.Product, error) {
	return queryEvent.repository.Find(name, category)
}
