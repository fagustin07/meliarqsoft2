package model

import (
	"meliarqsoft2/pkg/exceptions/model"
)

type Stock struct {
	Amount int
}

func NewStock(amount int) (*Stock, error) {
	if amount < 0 {
		return nil, model.InvalidStockError{}
	} else {
		return &Stock{Amount: amount}, nil
	}
}
