package action

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/application/command/action"
	"meliarqsoft2/internal/application/command/query"
	"meliarqsoft2/internal/purchase/domain"
	"meliarqsoft2/internal/purchase/domain/ports"
	"time"
)

type MakePurchaseEvent struct {
	repository         ports.IPurchaseRepository
	findProduct        *query.FindProduct
	existUser          *query.ExistUser
	manageProductStock *action.ManageProductStock
}

func NewMakePurchaseEvent(repo ports.IPurchaseRepository, findProd *query.FindProduct, existUser *query.ExistUser,
	manageStock *action.ManageProductStock) *MakePurchaseEvent {
	return &MakePurchaseEvent{
		repository:         repo,
		findProduct:        findProd,
		existUser:          existUser,
		manageProductStock: manageStock,
	}
}

func (actionEvent MakePurchaseEvent) Execute(IDProduct uuid.UUID, IDUser uuid.UUID, units int) (*domain.Purchase, error) {
	product, err := actionEvent.findProduct.Execute(IDProduct)
	if err != nil {
		return nil, err
	}

	if err := actionEvent.existUser.Execute(IDUser); err != nil {
		return nil, err
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("failed generating a purchase identifier")
	}

	newPurchase, err := domain.NewPurchase(newUUID, IDProduct, IDUser, time.Now(), units, product.Price.Value*float32(units))
	if err != nil {
		return nil, err
	}

	err = actionEvent.manageProductStock.Execute(IDProduct, units, false)
	if err != nil {
		return nil, err
	}

	err = actionEvent.repository.Create(newPurchase)
	if err != nil {
		err := actionEvent.manageProductStock.Execute(IDProduct, units, true)
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	return newPurchase, nil
}
