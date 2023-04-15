package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type FindProduct struct {
	repository domain.IProductRepository
}

func NewFindProductCommand(repository domain.IProductRepository) *FindProduct {
	return &FindProduct{repository: repository}
}

func (queryCommand FindProduct) Execute(productID uuid.UUID) (*domain.Product, error) {
	prod, err := queryCommand.repository.FindById(productID)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
