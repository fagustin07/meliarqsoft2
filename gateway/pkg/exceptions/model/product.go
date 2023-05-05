package model

import (
	"fmt"
)

type ProductNotFound struct{ Id string }

func (p ProductNotFound) Error() string {
	return fmt.Sprintf("cannot find product %v", p.Id)
}

type ProductAlreadyExist struct {
	Name     string
	SellerId string
}

func (p ProductAlreadyExist) Error() string {
	return fmt.Sprintf("product %v already exist for seller %v", p.Name, p.SellerId)
}

type InvalidPriceError struct{}

func (i InvalidPriceError) Error() string {
	return "invalid price value"
}

type InvalidStockError struct{}

func (i InvalidStockError) Error() string {
	return "invalid stock amount"
}

type ProductWithoutStockError struct{ ProductName string }

func (p ProductWithoutStockError) Error() string {
	return "cannot consume more units of " + p.ProductName + " than are in stock"
}

type ProductRestoreStockError struct {
	ProductName string
	Units       int
}

func (p ProductRestoreStockError) Error() string {
	return fmt.Sprintf("failed to restore %v units to product %v", p.Units, p.ProductName)
}

type MinAndMaxPriceCombinationError struct{}

func (m MinAndMaxPriceCombinationError) Error() string {
	return "min_price cannot be gt max_price"
}
