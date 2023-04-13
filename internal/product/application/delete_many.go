package application

import (
	"errors"
	"github.com/google/uuid"
)

func (service *ProductService) DeleteMany(sellerId uuid.UUID) error {
	_, err := service.repo.GetFrom(sellerId)
	if err != nil {
		return errors.New("failed getting all products")
	}

	return service.repo.DeleteMany(sellerId)
}
