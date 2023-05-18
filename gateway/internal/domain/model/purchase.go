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

type IFindPurchaseService interface {
	Execute(productID uuid.UUID) ([]PurchaseDTO, error)
}

type PurchaseDTO struct {
	ID        uuid.UUID `json:"id" bson:"id,omitempty"`
	IDProduct uuid.UUID `json:"id_product" bson:"id_product,omitempty"`
	IDUser    uuid.UUID `json:"id_user" bson:"id_user,omitempty"`
	Date      time.Time `json:"date" bson:"date"`
	Units     int       `json:"units" bson:"units"`
	Total     float32   `json:"total" bson:"total"`
}

func MapPurchaseToJSON(purchase *Purchase) PurchaseDTO {
	return PurchaseDTO{
		ID:        purchase.ID,
		IDProduct: purchase.IDProduct,
		IDUser:    purchase.IDUser,
		Date:      purchase.Date,
		Units:     purchase.Units.Amount,
		Total:     purchase.Total.Value,
	}
}
