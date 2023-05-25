package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type FindProductById struct {
	repository model.IProductRepository
}

func NewFindProductByIdEvent(repository model.IProductRepository) *FindProductById {
	return &FindProductById{repository: repository}
}

func (queryEvent FindProductById) Execute(productID uuid.UUID) (*model.Product, error) {
	prod, err := queryEvent.repository.FindById(productID)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
