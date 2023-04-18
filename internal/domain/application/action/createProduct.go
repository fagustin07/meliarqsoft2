package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/model"
)

type CreateProductEvent struct {
	productRepository model.IProductRepository
	existSeller       *query.ExistSeller
}

func NewCreateProductEvent(productRepository model.IProductRepository, existSeller *query.ExistSeller) *CreateProductEvent {
	return &CreateProductEvent{
		productRepository: productRepository,
		existSeller:       existSeller,
	}
}

func (actionEvent CreateProductEvent) Execute(product model.Product) (uuid.UUID, error) {
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
