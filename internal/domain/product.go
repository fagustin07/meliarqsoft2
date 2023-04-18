package domain

import (
	"errors"
	"github.com/google/uuid"
	"strconv"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Category    string
	Price       *Price
	Stock       *Stock
	IDSeller    uuid.UUID
}

func NewProduct(id uuid.UUID, name string, description string, category string, price float32, stock int, IDSeller uuid.UUID) (*Product, error) {
	priceObj, err := NewPrice(price)
	if err != nil {
		return nil, err
	}
	stockObj, err := NewStock(stock)
	if err != nil {
		return nil, err
	}

	return &Product{ID: id, Name: name, Description: description, Category: category, Price: priceObj, Stock: stockObj, IDSeller: IDSeller}, nil
}

func (prod *Product) CanConsume(units int) bool {
	return units+prod.Stock.Amount >= 0
}

func (prod *Product) TakeUnits(units int) (float32, int, error) {
	var unitsAbs float32
	newStock, err := NewStock(prod.Stock.Amount + units)
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

func (prod *Product) Consume(units int) error {
	newStock, err := NewStock(prod.Stock.Amount - units)
	if err != nil {
		return errors.New("cannot consume more units of " + prod.Name + " than are in stock")
	}
	prod.Stock = newStock
	return nil
}

func (prod *Product) Restore(units int) error {
	newStock, err := NewStock(prod.Stock.Amount + units)
	if err != nil {
		return errors.New("failed to restore " + strconv.Itoa(units) + "units to product " + prod.Name)
	}
	prod.Stock = newStock
	return nil
}

type IProductRepository interface {
	Create(product *Product) (*Product, error)
	Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (*Product, error)
	Delete(ID uuid.UUID) error
	Find(name string, category string) ([]*Product, error)
	Filter(minPrice float32, maxPrice float32) ([]*Product, error)
	FindById(ID uuid.UUID) (*Product, error)
	UpdateStock(ID uuid.UUID, stock int) error
	GetFrom(sellerId uuid.UUID) ([]*Product, error)
}
