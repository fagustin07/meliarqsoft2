package command

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/product/domain/ports"
)

type DeleteProduct struct {
	productRepository          ports.IProductRepository
	deletePurchasesByProductId DeletePurchasesByProductId
}

func NewDeleteProduct(prodRepo ports.IProductRepository, deletePurchases DeletePurchasesByProductId) *DeleteProduct {
	return &DeleteProduct{
		productRepository:          prodRepo,
		deletePurchasesByProductId: deletePurchases,
	}
}

func (usecase DeleteProduct) Execute(id uuid.UUID) error {
	if _, err := usecase.productRepository.GetProduct(id); err != nil {
		return err
	}

	if err := usecase.deletePurchasesByProductId.Execute(id); err != nil {
		return err
	}

	if err := usecase.productRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
