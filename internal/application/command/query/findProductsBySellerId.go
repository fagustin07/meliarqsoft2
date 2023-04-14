package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type FindProductsBySellerID struct {
	repository ports.IProductRepository
}

func (query FindProductsBySellerID) Execute(sellerID uuid.UUID) ([]*domain.Product, error) {
	res, err := query.repository.GetFrom(sellerID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
