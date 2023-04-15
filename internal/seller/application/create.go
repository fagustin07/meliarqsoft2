package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

func (service SellerService) Create(businessName string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	newSeller, err := domain.NewSeller(newUUID, businessName, email)
	if err != nil {
		return uuid.Nil, err
	}

	err = service.repo.Create(newSeller)
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}
