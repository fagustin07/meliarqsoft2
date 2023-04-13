package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/infrastructure/repository"
)

type FindProductsBySellerID struct {
	repository repository.ProductMongoDBRepository
}

func (query FindProductsBySellerID) Execute(sellerID uuid.UUID) ([]*domain.Product, error) {
	res, err := query.repository.GetFrom(sellerID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
