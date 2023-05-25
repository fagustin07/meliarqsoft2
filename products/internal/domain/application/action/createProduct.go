package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type CreateProductEvent struct {
	productRepository model.IProductRepository
	FindSellerById    model.IFindSellerByIdService
}

func NewCreateProductEvent(productRepository model.IProductRepository, findSeller model.IFindSellerByIdService) *CreateProductEvent {
	return &CreateProductEvent{
		productRepository: productRepository,
		FindSellerById:    findSeller,
	}
}

func (actionEvent CreateProductEvent) Execute(product model.Product) (uuid.UUID, error) {
	_, err := actionEvent.FindSellerById.Execute(product.IDSeller)

	if err != nil {
		return uuid.Nil, err
	}

	id, err := actionEvent.productRepository.Create(product)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
