package model

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

func (prod *Product) CanConsume(units int) bool {
	return prod.Stock.Amount >= units
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
		return errors.New("failed to restore " + strconv.Itoa(units) + " units to product " + prod.Name)
	}
	prod.Stock = newStock
	return nil
}

//go:generate mockgen -destination=../mock/productRepository.go -package=mock -source=product.go
type IProductRepository interface {
	Create(product Product) (uuid.UUID, error)
	Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (Product, error)
	Delete(ID uuid.UUID) error
	Find(name string, category string) ([]Product, error)
	Filter(minPrice float32, maxPrice float32) ([]Product, error)
	FindById(ID uuid.UUID) (*Product, error)
	UpdateStock(ID uuid.UUID, stock int) error
	GetFrom(sellerId uuid.UUID) ([]Product, error)
}
