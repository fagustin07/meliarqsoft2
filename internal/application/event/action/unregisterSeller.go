package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/command/action"
	"meliarqsoft2/internal/application/command/query"
	"meliarqsoft2/internal/domain"
)

type UnregisterSellerEvent struct {
	sellerRepository       domain.ISellerRepository
	deleteProductsBySeller *action.DeleteProductsBySeller
	existSeller            *query.ExistSeller
}

func NewUnregisterSellerEvent(sellerRepository domain.ISellerRepository,
	deleteProductsBySeller *action.DeleteProductsBySeller, existSeller *query.ExistSeller) *UnregisterSellerEvent {
	return &UnregisterSellerEvent{
		sellerRepository:       sellerRepository,
		deleteProductsBySeller: deleteProductsBySeller,
		existSeller:            existSeller,
	}
}

func (actionEvent UnregisterSellerEvent) Execute(id uuid.UUID) error {
	if err := actionEvent.existSeller.Execute(id); err != nil {
		return err
	}

	if err := actionEvent.deleteProductsBySeller.Execute(id); err != nil {
		return err
	}

	if err := actionEvent.sellerRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
