package application

import (
	"errors"
	"github.com/google/uuid"
)

func (service *ProductApplication) DeleteMany(sellerId uuid.UUID) error {
	products, err := service.repo.GetFrom(sellerId)
	if err != nil {
		return errors.New("failed getting all products")
	}
	for _, prod := range products {
		err := service.purchaseService.DeleteMany(prod.ID)
		if err != nil {
			return errors.New("failed deleting all products purchases")
		}
	}

	return service.repo.DeleteMany(sellerId)
}
