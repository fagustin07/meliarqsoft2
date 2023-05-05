package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
)

type UndoPurchasesByProduct struct {
	findPurchasesFromProduct *query.FindPurchasesFromProductEvent
	undoPurchase             *UndoPurchase
}

func NewUndoPurchasesByProductCommand(findPurchase *query.FindPurchasesFromProductEvent,
	undoPurchase *UndoPurchase) *UndoPurchasesByProduct {

	return &UndoPurchasesByProduct{
		findPurchasesFromProduct: findPurchase,
		undoPurchase:             undoPurchase,
	}
}

func (commandAction UndoPurchasesByProduct) Execute(productID uuid.UUID) error {
	purchases, err := commandAction.findPurchasesFromProduct.Execute(productID)
	if err != nil {
		return err
	}

	for _, purchase := range purchases {
		err := commandAction.undoPurchase.Execute(purchase.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
