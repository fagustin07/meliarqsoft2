package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type FindProductsBySeller struct {
	repository ports.IProductRepository
}

func NewFindProductsBySellerCommand(repository ports.IProductRepository) *FindProductsBySeller {
	return &FindProductsBySeller{repository: repository}
}

func (queryCommand FindProductsBySeller) Execute(sellerID uuid.UUID) ([]*domain.Product, error) {
	res, err := queryCommand.repository.GetFrom(sellerID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
