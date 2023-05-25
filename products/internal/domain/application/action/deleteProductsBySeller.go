package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type DeleteProductsBySeller struct {
	productRepository      model.IProductRepository
	deletePurchasesService model.IDeletePurchasesByProductsService
}

func NewDeleteProductsBySellerEvent(productRepo model.IProductRepository, deletePurchases model.IDeletePurchasesByProductsService) *DeleteProductsBySeller {
	return &DeleteProductsBySeller{
		productRepository:      productRepo,
		deletePurchasesService: deletePurchases,
	}
}

func (actionEvent DeleteProductsBySeller) Execute(sellerID uuid.UUID) ([]uuid.UUID, error) {
	deletedIds, err := actionEvent.productRepository.DeleteBySeller(sellerID)
	if err != nil {
		return nil, err
	}

	deletePurchErr := actionEvent.deletePurchasesService.Execute(deletedIds)

	if deletePurchErr != nil { // si falla borrado, comienzo rollback
		_, err := actionEvent.productRepository.Restore(deletedIds)
		if err != nil {
			// se le avisa a backoffice el error para restauración manual
			return nil, err
		}

		return nil, deletePurchErr
	}

	return deletedIds, nil
}
