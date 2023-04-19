package model

import (
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

//go:generate mockgen -destination=../mock/purchaseRepository.go -package=mock -source=purchase.go
type IPurchaseRepository interface {
	Create(purchase *Purchase, product *Product) (uuid.UUID, float32, error)
	Find(productID uuid.UUID) ([]*Purchase, error)
	Delete(ID uuid.UUID) error
}
