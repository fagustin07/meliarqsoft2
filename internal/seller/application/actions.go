package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain"
)

func (manager SellerManager) Create(businessName string, email string) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return newUUID, err
	}

	newSeller := domain.NewSeller(newUUID, businessName, email)
	err = manager.repo.Create(newSeller)
	if err != nil {
		return newUUID, err
	}

	return newUUID, nil
}

func (manager SellerManager) Update(id uuid.UUID, businessName string, email string) error {
	return manager.repo.Update(id, businessName, email)
}

func (manager SellerManager) Delete(ID uuid.UUID) error {
	return manager.repo.Delete(ID)
}

func (manager SellerManager) Find(businessName string) ([]*domain.Seller, error) {
	return manager.repo.Find(businessName)
}
