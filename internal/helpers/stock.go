package helpers

import "errors"

type Stock struct {
	Amount int
}

func NewStock(amount int) (*Stock, error) {
	if amount < 0 {
		return &Stock{}, errors.New("invalid stock amount")
	} else {
		return &Stock{Amount: amount}, nil
	}
}
