package query

import (
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type FindProduct struct {
	repository ports.IProductRepository
}

func (queryEvent FindProduct) Execute(name string, category string) ([]*domain.Product, error) {
	return queryEvent.repository.Find(name, category)
}
