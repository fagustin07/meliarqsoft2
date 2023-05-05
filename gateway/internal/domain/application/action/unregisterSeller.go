package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/model"
)

type UnregisterSellerEvent struct {
	sellerRepository       model.ISellerRepository
	deleteProductsBySeller *DeleteProductsBySeller
	existSeller            *query.ExistSeller
}

func NewUnregisterSellerEvent(sellerRepository model.ISellerRepository,
	deleteProductsBySeller *DeleteProductsBySeller, existSeller *query.ExistSeller) *UnregisterSellerEvent {
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
