package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain"
)

func (service SellerService) Create(businessName string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return newUUID, err
	}

	newSeller := domain.NewSeller(newUUID, businessName, email)
	err = service.repo.Create(newSeller)
	if err != nil {
		return newUUID, err
	}

	return newUUID, nil
}
