package command

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain/ports"
)

type DeletePurchasesByProductId struct {
	repository     ports.IPurchaseRepository
	deletePurchase DeletePurchase
}

func NewDeletePurchasesByProductId(repository ports.IPurchaseRepository, deletePurchase DeletePurchase) *DeletePurchasesByProductId {
	return &DeletePurchasesByProductId{
		repository:     repository,
		deletePurchase: deletePurchase,
	}
}

func (command DeletePurchasesByProductId) Execute(productID uuid.UUID) error {
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
	err = command.deletePurchase.Execute(productID)
	if err != nil {
		return err
	}

	return nil
}
