package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type RegisterSellerEvent struct {
	repository model.ISellerRepository
}

func NewRegisterSellerEvent(repository model.ISellerRepository) *RegisterSellerEvent {
	return &RegisterSellerEvent{repository: repository}
}

func (event RegisterSellerEvent) Execute(seller *model.Seller) (uuid.UUID, error) {
	newId, err := event.repository.Create(seller)
	if err != nil {
		return uuid.Nil, err
	}

	return newId, nil
}
