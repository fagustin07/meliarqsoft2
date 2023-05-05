package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/model"
)

type MakePurchaseEvent struct {
	repository         model.IPurchaseRepository
	findProduct        *query.FindProductById
	existUser          *query.ExistUser
	manageProductStock *ManageProductStock
}

func NewMakePurchaseEvent(repo model.IPurchaseRepository, findProd *query.FindProductById, existUser *query.ExistUser,
	manageStock *ManageProductStock) *MakePurchaseEvent {
	return &MakePurchaseEvent{
		repository:         repo,
		findProduct:        findProd,
		existUser:          existUser,
		manageProductStock: manageStock,
	}
}

func (actionEvent MakePurchaseEvent) Execute(purchase *model.Purchase) (uuid.UUID, float32, error) {
	product, err := actionEvent.findProduct.Execute(purchase.IDProduct)
	if err != nil {
		return uuid.Nil, 0, err
	}

	if err := actionEvent.existUser.Execute(purchase.IDUser); err != nil {
		return uuid.Nil, 0, err
	}

	err = actionEvent.manageProductStock.Execute(purchase.IDProduct, purchase.Units.Amount, false)
	if err != nil {
		return uuid.Nil, 0, err
	}

	id, total, err := actionEvent.repository.Create(purchase, product)

	if err != nil {
		err2 := actionEvent.manageProductStock.Execute(purchase.IDProduct, purchase.Units.Amount, true)
		if err2 != nil {
			return uuid.Nil, 0, err2
		}

		return uuid.Nil, 0, err
	}

	return id, total, nil
}
