package application

import (
	"errors"
	"github.com/google/uuid"
)

func (service *ProductApplication) Delete(ID uuid.UUID) error {
	err := service.purchaseService.DeleteMany(ID)
	if err != nil {
		return errors.New("failed to delete all purchases from product " + ID.String())
	}
	return service.repo.Delete(ID)
}
