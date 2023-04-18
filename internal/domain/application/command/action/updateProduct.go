package action

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type UpdateProductEvent struct {
	repository domain.IProductRepository
}

func NewUpdateProductEvent(repository domain.IProductRepository) *UpdateProductEvent {
	return &UpdateProductEvent{
		repository: repository,
	}
}

func (actionEvent UpdateProductEvent) Execute(ID uuid.UUID, name string, description string, category string, price float32, stock int) (domain.Product, error) {
	newPrice, err := domain.NewPrice(price)
	if err != nil {
		return domain.Product{}, err
	}
	newStock, err := domain.NewStock(stock)
	if err != nil {
		return domain.Product{}, err
	}

	return actionEvent.repository.Update(ID, name, description, category, newPrice.Value, newStock.Amount)
}
