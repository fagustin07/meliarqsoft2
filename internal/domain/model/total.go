package model

import "errors"

type Total struct {
	Value float32
}

func NewTotal(value float32) (*Total, error) {
	if value < 0 {
		return nil, errors.New("invalid units amount")
	}

	return &Total{Value: value}, nil
}
