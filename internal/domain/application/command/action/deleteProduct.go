package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type DeleteProductEvent struct {
	productRepository      domain.IProductRepository
	undoPurchasesByProduct *UndoPurchasesByProduct
}

func NewDeleteProductEvent(prodRepo domain.IProductRepository, undoPurchasesByProduct *UndoPurchasesByProduct) *DeleteProductEvent {
	return &DeleteProductEvent{
		productRepository:      prodRepo,
		undoPurchasesByProduct: undoPurchasesByProduct,
	}
}

func (actionEvent DeleteProductEvent) Execute(id uuid.UUID) error {
	if _, err := actionEvent.productRepository.FindById(id); err != nil {
		return err
	}

	if err := actionEvent.undoPurchasesByProduct.Execute(id); err != nil {
		return err
	}

	if err := actionEvent.productRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
