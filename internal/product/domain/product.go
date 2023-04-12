package domain

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/helpers"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Category    string
	Price       *helpers.Price
	Stock       *helpers.Stock
	IDSeller    uuid.UUID
}

func NewProduct(id uuid.UUID, name string, description string, category string, price float32, stock int, IDSeller uuid.UUID) (*Product, error) {
	priceObj, err := helpers.NewPrice(price)
	if err != nil {
		return nil, err
	}
	stockObj, err := helpers.NewStock(stock)
	if err != nil {
		return nil, err
	}

	return &Product{ID: id, Name: name, Description: description, Category: category, Price: priceObj, Stock: stockObj, IDSeller: IDSeller}, nil
}

func (prod Product) CanConsume(units int) bool {
	return units+prod.Stock.Amount >= 0
}

func (prod Product) TakeUnits(units int) (float32, int, error) {
	var unitsAbs float32
	newStock, err := helpers.NewStock(prod.Stock.Amount + units)
	if err != nil {
		return 0, 0, err
	}

	if units < 0 {
		unitsAbs = float32(units * -1)
	} else {
		unitsAbs = float32(units)
	}

	prod.Stock = newStock

	return prod.Price.Value * unitsAbs, prod.Stock.Amount, nil
}
