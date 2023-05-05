package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
)

type UndoPurchasesByProductEvent struct {
	findPurchasesFromProduct *query.FindPurchasesFromProductEvent
	undoPurchase             *UndoPurchase
}

func NewUndoPurchasesByProductEvent(findPurchase *query.FindPurchasesFromProductEvent,
	undoPurchase *UndoPurchase) *UndoPurchasesByProductEvent {

	return &UndoPurchasesByProductEvent{
		findPurchasesFromProduct: findPurchase,
		undoPurchase:             undoPurchase,
	}
}

func (actionEvent UndoPurchasesByProductEvent) Execute(productID uuid.UUID) error {
	purchases, err := actionEvent.findPurchasesFromProduct.Execute(productID)
	if err != nil {
		return err
	}

	for _, purchase := range purchases {
		err := actionEvent.undoPurchase.Execute(purchase.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
