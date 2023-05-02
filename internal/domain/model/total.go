package model

import (
	"meliarqsoft2/pkg/exceptions/model"
)

type Total struct {
	Value float32
}

func NewTotal(value float32) (*Total, error) {
	if value < 0 {
		return nil, model.InvalidTotalPriceError{}
	}

	return &Total{Value: value}, nil
}
