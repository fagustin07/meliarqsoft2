package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type FindProductById struct {
	repository domain.IProductRepository
}

func NewFindProductCommand(repository domain.IProductRepository) *FindProductById {
	return &FindProductById{repository: repository}
}

func (queryCommand FindProductById) Execute(productID uuid.UUID) (domain.Product, error) {
	prod, err := queryCommand.repository.FindById(productID)
	if err != nil {
		return domain.Product{}, err
	}

	return prod, nil
}
