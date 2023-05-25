package model

import (
	"meliarqsoft2/pkg/exceptions/model"
)

type Units struct {
	Amount int
}

func NewUnits(value int) (*Units, error) {
	if value <= 0 {
		return nil, model.InvalidUnitsError{}
	}
	return &Units{Amount: value}, nil
}
