package model

import "errors"

type Units struct {
	Amount int
}

func NewUnits(value int) (*Units, error) {
	if value <= 0 {
		return nil, errors.New("invalid units amount")
	}
	return &Units{Amount: value}, nil
}
