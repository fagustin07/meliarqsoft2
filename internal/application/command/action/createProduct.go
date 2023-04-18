package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/command/query"
	"meliarqsoft2/internal/domain"
)

type CreateProductEvent struct {
	productRepository domain.IProductRepository
	existSeller       *query.ExistSeller
}

func NewCreateProductEvent(productRepository domain.IProductRepository, existSeller *query.ExistSeller) *CreateProductEvent {
	return &CreateProductEvent{
		productRepository: productRepository,
		existSeller:       existSeller,
	}
}

func (actionEvent CreateProductEvent) Execute(product domain.Product) (uuid.UUID, error) {
	err := actionEvent.existSeller.Execute(product.IDSeller)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := actionEvent.productRepository.Create(product)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
