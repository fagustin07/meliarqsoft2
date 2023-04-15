package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain"
	"meliarqsoft2/internal/seller/domain/ports"
)

type RegisterSellerEvent struct {
	repository ports.ISellerRepository
}

func NewRegisterSellerEvent(repository ports.ISellerRepository) *RegisterSellerEvent {
	return &RegisterSellerEvent{repository: repository}
}

func (event RegisterSellerEvent) Execute(businessName string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	newSeller, err := domain.NewSeller(newUUID, businessName, email)
	if err != nil {
		return uuid.Nil, err
	}

	err = event.repository.Create(newSeller)
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}
