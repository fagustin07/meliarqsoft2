package application

import (
	"errors"
	"github.com/google/uuid"
)

func (service *ProductApplication) ManipulateStock(idProduct uuid.UUID, units int) (float32, error) {
	prod, err := service.repo.GetProduct(idProduct)
	if err != nil {
		return 0, err
	}

	if !prod.CanConsume(units) {
		return 0, errors.New("cannot buy more units than are in stock")
	} else {
		totalPrice, currentStock := prod.TakeUnits(units)
		err := service.repo.UpdateStock(idProduct, currentStock)
		if err != nil {
			return 0, err
		}

		return totalPrice, nil
	}
}
