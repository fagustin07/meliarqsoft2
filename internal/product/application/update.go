package application

import (
	"github.com/google/uuid"
	domain2 "meliarqsoft2/internal/domain"
	"meliarqsoft2/internal/product/domain"
)

func (service *ProductService) Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (*domain.Product, error) {
	newPrice, err := domain2.NewPrice(price)
	if err != nil {
		return nil, err
	}
	newStock, err := domain2.NewStock(stock)
	if err != nil {
		return nil, err
	}

	return service.repo.Update(ID, name, description, category, newPrice.Value, newStock.Amount)
}
