package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/command/query"
)

type UndoPurchasesByProduct struct {
	findPurchases *query.FindPurchasesFromProduct
	undoPurchase  *DeletePurchase
}

func NewUndoPurchasesByProductCommand(findPurchase *query.FindPurchasesFromProduct, undoPurchase *DeletePurchase) *UndoPurchasesByProduct {
	return &UndoPurchasesByProduct{
		findPurchases: findPurchase,
		undoPurchase:  undoPurchase,
	}
}

func (commandAction UndoPurchasesByProduct) Execute(productID uuid.UUID) error {
	purchases, err := commandAction.findPurchases.Execute(productID)
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
