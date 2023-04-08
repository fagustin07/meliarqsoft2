package application

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/purchase/domain"
	"time"
)

func (manager PurchaseManager) Create(IDProduct uuid.UUID, IDUser uuid.UUID, units int) (*domain.Purchase, error) {
	if units == 0 {
		return &domain.Purchase{}, errors.New("cannot buy zero units")
	}
	err := manager.userManager.Exist(IDUser)
	if err != nil {
		return &domain.Purchase{}, err
	}
	purchaseTotal, err := manager.productManager.ManipulateStock(IDProduct, units*-1)
	if err != nil {
		return &domain.Purchase{}, err
	}

	newPurchase := domain.NewPurchase(IDProduct, IDUser, time.Now(), units, purchaseTotal)
	err = manager.repository.Create(newPurchase)
	if err != nil {
		_, err := manager.productManager.ManipulateStock(IDProduct, units)
		if err != nil {
			return nil, err
		}
		return &domain.Purchase{}, err
	}

	return newPurchase, nil
}
