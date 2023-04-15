package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type FindProductsBySeller struct {
	repository domain.IProductRepository
}

func NewFindProductsBySellerCommand(repository domain.IProductRepository) *FindProductsBySeller {
	return &FindProductsBySeller{repository: repository}
}

func (queryCommand FindProductsBySeller) Execute(sellerID uuid.UUID) ([]*domain.Product, error) {
	res, err := queryCommand.repository.GetFrom(sellerID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
