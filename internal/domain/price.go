package domain

import "errors"

type Price struct {
	Value float32
}

func NewPrice(amount float32) (*Price, error) {
	if amount < 0 {
		return &Price{}, errors.New("invalid price value")
	} else {
		return &Price{Value: amount}, nil
	}
}
