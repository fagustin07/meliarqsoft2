package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type FindProduct struct {
	repository ports.IProductRepository
}

func NewFindProductCommand(repository ports.IProductRepository) *FindProduct {
	return &FindProduct{repository: repository}
}

func (queryCommand FindProduct) Execute(productID uuid.UUID) (*domain.Product, error) {
	prod, err := queryCommand.repository.FindById(productID)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
