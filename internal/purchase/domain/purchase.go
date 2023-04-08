package domain

import (
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	IDProduct uuid.UUID
	IDUser    uuid.UUID
	Date      time.Time
	Units     int
	Total     float32
}

func NewPurchase(IDProduct uuid.UUID, IDUser uuid.UUID, date time.Time, units int, total float32) *Purchase {
	return &Purchase{IDProduct: IDProduct, IDUser: IDUser, Date: date, Units: units, Total: total}
}
