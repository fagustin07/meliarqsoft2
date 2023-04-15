package domain

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
	"time"
)

type Purchase struct {
	ID        uuid.UUID
	IDProduct uuid.UUID
	IDUser    uuid.UUID
	Date      time.Time
	Units     *domain.Units
	Total     *domain.Total
}

func NewPurchase(ID uuid.UUID, IDProduct uuid.UUID, IDUser uuid.UUID, date time.Time, units int, total float32) (*Purchase, error) {
	newUnits, err1 := domain.NewUnits(units)
	newTotal, err2 := domain.NewTotal(total)
	if err1 != nil || err2 != nil {
		return nil, errors.New("failed generating new purchase units or total price")
	}

	return &Purchase{ID: ID, IDProduct: IDProduct, IDUser: IDUser, Date: date, Units: newUnits, Total: newTotal}, nil
}
