package query

import (
	"meliarqsoft2/internal/domain"
)

type FindProductEvent struct {
	repository domain.IProductRepository
}

func NewFindProductEvent(repository domain.IProductRepository) *FindProductEvent {
	return &FindProductEvent{
		repository: repository,
	}
}

func (queryEvent FindProductEvent) Execute(name string, category string) ([]*domain.Product, error) {
	return queryEvent.repository.Find(name, category)
}
