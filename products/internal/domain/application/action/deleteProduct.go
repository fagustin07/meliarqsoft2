package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type DeleteProductEvent struct {
	productRepository      model.IProductRepository
	deletePurchasesService model.IDeletePurchasesByProductsService
}

func NewDeleteProductEvent(prodRepo model.IProductRepository, deletePurchases model.IDeletePurchasesByProductsService) *DeleteProductEvent {
	return &DeleteProductEvent{
		productRepository:      prodRepo,
		deletePurchasesService: deletePurchases,
	}
}

func (actionEvent DeleteProductEvent) Execute(id uuid.UUID) error {
	err := actionEvent.productRepository.Delete(id)

	if err != nil {
		return err
	}

	deletePurchErr := actionEvent.deletePurchasesService.Execute([]uuid.UUID{id})

	if deletePurchErr != nil { // si falla borrado de compras, rollback de producto
		_, err2 := actionEvent.productRepository.Restore([]uuid.UUID{id})
		if err2 != nil {
			// se le avisa a backoffice el error para restauración manual
			return err
		}

		return deletePurchErr
	}
	return nil
}
