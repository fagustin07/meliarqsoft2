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

func (event RegisterSellerEvent) Execute(businessName string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	newSeller, err := model.NewSeller(newUUID, businessName, email)
	if err != nil {
		return uuid.Nil, err
	}

	err = event.repository.Create(newSeller)
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}
