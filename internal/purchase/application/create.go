package application

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
	"time"
)

func (service PurchaseService) Create(IDProduct uuid.UUID, IDUser uuid.UUID, units int) (*domain.Purchase, error) {
	if units == 0 {
		return &domain.Purchase{}, errors.New("cannot buy zero units")
	}
	err := service.userService.Exist(IDUser)
	if err != nil {
		return &domain.Purchase{}, err
	}
	purchaseTotal, err := service.productService.ManipulateStock(IDProduct, units*-1)
	if err != nil {
		return &domain.Purchase{}, err
	}

	newPurchase := domain.NewPurchase(IDProduct, IDUser, time.Now(), units, purchaseTotal)
	err = service.repository.Create(newPurchase)
	if err != nil {
		_, err := service.productService.ManipulateStock(IDProduct, units)
		if err != nil {
			return nil, err
		}
		return &domain.Purchase{}, err
	}

	return newPurchase, nil
}
