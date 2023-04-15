package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	ID        uuid.UUID
	IDProduct uuid.UUID
	IDUser    uuid.UUID
	Date      time.Time
	Units     *Units
	Total     *Total
}

func NewPurchase(ID uuid.UUID, IDProduct uuid.UUID, IDUser uuid.UUID, date time.Time, units int, total float32) (*Purchase, error) {
	newUnits, err1 := NewUnits(units)
	newTotal, err2 := NewTotal(total)
	if err1 != nil || err2 != nil {
		return nil, errors.New("failed generating new purchase units or total price")
	}

	return &Purchase{ID: ID, IDProduct: IDProduct, IDUser: IDUser, Date: date, Units: newUnits, Total: newTotal}, nil
}

type IPurchaseRepository interface {
	Create(purchase *Purchase) error
	Find(productID uuid.UUID) ([]*Purchase, error)
	DeleteMany(productId uuid.UUID) error
	Delete(ID uuid.UUID) error
}
