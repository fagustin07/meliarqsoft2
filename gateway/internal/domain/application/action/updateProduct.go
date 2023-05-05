package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UpdateProductEvent struct {
	repository model.IProductRepository
}

func NewUpdateProductEvent(repository model.IProductRepository) *UpdateProductEvent {
	return &UpdateProductEvent{
		repository: repository,
	}
}

func (actionEvent UpdateProductEvent) Execute(ID uuid.UUID, name string, description string, category string, price float32, stock int) (model.Product, error) {
	newPrice, err := model.NewPrice(price)
	if err != nil {
		return model.Product{}, err
	}
	newStock, err := model.NewStock(stock)
	if err != nil {
		return model.Product{}, err
	}

	return actionEvent.repository.Update(ID, name, description, category, newPrice.Value, newStock.Amount)
}
