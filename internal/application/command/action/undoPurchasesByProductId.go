package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain/ports"
)

type UndoPurchasesByProductId struct {
	repository   ports.IPurchaseRepository
	undoPurchase *DeletePurchase
}

func NewUndoPurchasesByProductIdCommand(repository ports.IPurchaseRepository, undoPurchase *DeletePurchase) *UndoPurchasesByProductId {
	return &UndoPurchasesByProductId{
		repository:   repository,
		undoPurchase: undoPurchase,
	}
}

func (command UndoPurchasesByProductId) Execute(productID uuid.UUID) error {
	_, err := command.repository.Find(productID)
	if err != nil {
		return err
	}

	//TODO : Abstraer a command
	//  for _, purchase := range purchases {
	//	err := command.repository.Execute(purchase.ID)
	//	if err != nil {
	//		return err
	//	}
	//}
	err = command.undoPurchase.Execute(productID)
	if err != nil {
		return err
	}

	return nil
}
