package application

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/helpers"
	"meliarqsoft2/internal/product/domain"
)

func (service *ProductApplication) Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (*domain.Product, error) {
	newPrice, err := helpers.NewPrice(price)
	if err != nil {
		return nil, err
	}
	newStock, err := helpers.NewStock(stock)
	if err != nil {
		return nil, err
	}

	return service.repo.Update(ID, name, description, category, newPrice.Value, newStock.Amount)
}
