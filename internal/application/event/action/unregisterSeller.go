package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/command/action"
	"meliarqsoft2/internal/seller/domain/ports"
)

type UnregisterSellerEvent struct {
	sellerRepository         ports.ISellerRepository
	deleteProductsBySellerId *action.DeleteProductsBySeller
}

func NewDeleteSellerEvent(sellerRepo ports.ISellerRepository, deleteProd *action.DeleteProductsBySeller) *UnregisterSellerEvent {
	return &UnregisterSellerEvent{
		sellerRepository:         sellerRepo,
		deleteProductsBySellerId: deleteProd,
	}
}

func (usecase UnregisterSellerEvent) Execute(id uuid.UUID) error {
	if err := usecase.sellerRepository.Exist(id); err != nil {
		return err
	}

	if err := usecase.deleteProductsBySellerId.Execute(id); err != nil {
		return err
	}

	if err := usecase.sellerRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
