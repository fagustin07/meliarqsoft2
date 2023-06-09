package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type ManageProductStock struct {
	repository model.IProductRepository
}

func NewManageProductStockCommand(repository model.IProductRepository) *ManageProductStock {
	return &ManageProductStock{repository: repository}
}

func (actionCommand ManageProductStock) Execute(ID uuid.UUID, units int, isIncrement bool) error {
	product, err := actionCommand.repository.FindById(ID)
	if err != nil {
		return err
	}

	if !isIncrement {
		err := product.Consume(units)
		if err != nil {
			return err
		}
	} else {
		err := product.Restore(units)
		if err != nil {
			return err
		}
	}

	err = actionCommand.repository.UpdateStock(ID, product.Stock.Amount)
	if err != nil {
		return err
	}

	return nil
}
