package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type FindProductsBySeller struct {
	repository model.IProductRepository
}

func NewFindProductsBySellerCommand(repository model.IProductRepository) *FindProductsBySeller {
	return &FindProductsBySeller{repository: repository}
}

func (queryCommand FindProductsBySeller) Execute(sellerID uuid.UUID) ([]model.Product, error) {
	res, err := queryCommand.repository.GetFrom(sellerID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
