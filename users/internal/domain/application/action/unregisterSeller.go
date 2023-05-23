package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UnregisterSellerEvent struct {
	sellerRepository      model.ISellerRepository
	DeleteProductsService model.IDeleteProductsBySellerService
}

func NewUnregisterSellerEvent(sellerRepository model.ISellerRepository, deleteProdService model.IDeleteProductsBySellerService) *UnregisterSellerEvent {
	return &UnregisterSellerEvent{
		sellerRepository:      sellerRepository,
		DeleteProductsService: deleteProdService,
	}
}

func (actionEvent UnregisterSellerEvent) Execute(id uuid.UUID) error {
	err := actionEvent.sellerRepository.Delete(id)
	if err != nil {
		return err
	}

	failErr := actionEvent.DeleteProductsService.Execute(id)

	if failErr != nil { // rollback vendedor
		err2 := actionEvent.sellerRepository.Restore(id)
		if err2 != nil {
			// notificar a un back office el error para restauracion manual
			return err
		}

		return failErr
	}

	return nil
}
