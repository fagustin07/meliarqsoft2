package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type ManageProductStock struct {
	repository domain.IProductRepository
}

func NewManageProductStockCommand(repository domain.IProductRepository) *ManageProductStock {
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
