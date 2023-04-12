package domain

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/helpers/value_objects"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Category    string
	Price       *value_objects.Price
	Stock       *value_objects.Stock
	IDSeller    uuid.UUID
}

func NewProduct(id uuid.UUID, name string, description string, category string, price float32, stock int, IDSeller uuid.UUID) (*Product, error) {
	priceObj, err := value_objects.NewPrice(price)
	if err != nil {
		return nil, err
	}
	stockObj, err := value_objects.NewStock(stock)
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
	newStock, err := value_objects.NewStock(prod.Stock.Amount + units)
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
