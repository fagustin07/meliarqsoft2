package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UndoPurchasesByProductEvent struct {
	purchaseRepository model.IPurchaseRepository
}

func NewUndoPurchasesByProductEvent(purchaseRepository model.IPurchaseRepository) *UndoPurchasesByProductEvent {
	return &UndoPurchasesByProductEvent{
		purchaseRepository: purchaseRepository,
	}
}

func (actionEvent UndoPurchasesByProductEvent) Execute(productsIDs []uuid.UUID) error {
	err := actionEvent.purchaseRepository.DeleteByProductsIDs(productsIDs)

	if err != nil {
		return err
	}

	return nil
}
