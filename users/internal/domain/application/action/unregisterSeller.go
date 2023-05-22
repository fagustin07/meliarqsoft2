package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UnregisterSellerEvent struct {
	sellerRepository model.ISellerRepository
}

func NewUnregisterSellerEvent(sellerRepository model.ISellerRepository) *UnregisterSellerEvent {
	return &UnregisterSellerEvent{
		sellerRepository: sellerRepository,
	}
}

func (actionEvent UnregisterSellerEvent) Execute(id uuid.UUID) error {
	err := actionEvent.sellerRepository.Delete(id)
	if err != nil {
		return err
	}

	// TODO: notificar al servicio de productos que se elimino al vendedor. si falla, restaurar vendedor
	falloBorradoProds, err := err, nil

	if falloBorradoProds != nil { // rollback vendedor
		err2 := actionEvent.sellerRepository.Restore(id)
		if err2 != nil {
			// notificar a un back office el error para restauracion manual
			return err
		}

		return err // envia el motivo por el que ha fallado la coreografia
	}

	return nil
}
