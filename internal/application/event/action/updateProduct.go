package action

import (
	"github.com/google/uuid"
	domain2 "meliarqsoft2/internal/domain"
	"meliarqsoft2/internal/product/domain"
	"meliarqsoft2/internal/product/domain/ports"
)

type UpdateProductEvent struct {
	repository ports.IProductRepository
}

func NewUpdateProductEvent(repository ports.IProductRepository) *UpdateProductEvent {
	return &UpdateProductEvent{
		repository: repository,
	}
}

func (actionEvent UpdateProductEvent) Execute(ID uuid.UUID, name string, description string, category string, price float32, stock int) (*domain.Product, error) {
	newPrice, err := domain2.NewPrice(price)
	if err != nil {
		return nil, err
	}
	newStock, err := domain2.NewStock(stock)
	if err != nil {
		return nil, err
	}

	return actionEvent.repository.Update(ID, name, description, category, newPrice.Value, newStock.Amount)
}
