package model

import (
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	ID         uuid.UUID
	IDProduct  uuid.UUID
	IDCustomer uuid.UUID
	Date       time.Time
	Units      *Units
	Total      *Total
}

//go:generate mockgen -destination=../mock/purchaseRepository.go -package=mock -source=purchase.go
type IPurchaseRepository interface {
	Create(purchase *Purchase) (uuid.UUID, error)
	Find(productID uuid.UUID) ([]*Purchase, error)
	DeleteByProductsIDs(productsIDs []uuid.UUID) error
}
