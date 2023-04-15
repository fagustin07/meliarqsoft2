package query

import (
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type FindProductEvent struct {
	repository ports.IProductRepository
}

func NewFindProductEvent(repository ports.IProductRepository) *FindProductEvent {
	return &FindProductEvent{
		repository: repository,
	}
}

func (queryEvent FindProductEvent) Execute(name string, category string) ([]*domain.Product, error) {
	return queryEvent.repository.Find(name, category)
}
