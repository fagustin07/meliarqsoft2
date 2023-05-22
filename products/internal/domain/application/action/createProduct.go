package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type CreateProductEvent struct {
	productRepository model.IProductRepository
}

func NewCreateProductEvent(productRepository model.IProductRepository) *CreateProductEvent {
	return &CreateProductEvent{
		productRepository: productRepository,
	}
}

func (actionEvent CreateProductEvent) Execute(product model.Product) (uuid.UUID, error) {
	// TODO: consultar si existe el seller en el user service

	id, err := actionEvent.productRepository.Create(product)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
