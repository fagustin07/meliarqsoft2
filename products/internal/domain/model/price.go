package model

import (
	"meliarqsoft2/pkg/exceptions/model"
)

type Price struct {
	Value float32
}

func NewPrice(amount float32) (*Price, error) {
	if amount < 0 {
		return nil, model.InvalidPriceError{}
	} else {
		return &Price{Value: amount}, nil
	}
}
