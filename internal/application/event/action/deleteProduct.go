package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/command/action"
	"meliarqsoft2/internal/product/domain/ports"
)

type DeleteProductEvent struct {
	productRepository        ports.IProductRepository
	deletePurchasesByProduct *action.UndoPurchasesByProduct
}

func NewDeleteProductEvent(prodRepo ports.IProductRepository, deletePurchases *action.UndoPurchasesByProduct) *DeleteProductEvent {
	return &DeleteProductEvent{
		productRepository:        prodRepo,
		deletePurchasesByProduct: deletePurchases,
	}
}

func (usecase DeleteProductEvent) Execute(id uuid.UUID) error {
	if _, err := usecase.productRepository.GetProduct(id); err != nil {
		return err
	}

	if err := usecase.deletePurchasesByProduct.Execute(id); err != nil {
		return err
	}

	if err := usecase.productRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
