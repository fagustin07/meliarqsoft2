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
	Units     int
	Total     float32
}

type CreatePurchaseRequest struct {
	IDProduct uuid.UUID `json:"product_id" bson:"id_product,omitempty"`
	IDUser    uuid.UUID `json:"customer_id" bson:"id_user,omitempty"`
	Units     int       `json:"units" bson:"units"`
}

type IFindPurchaseService interface {
	Execute(productID uuid.UUID) ([]Purchase, error)
}

type IMakePurchaseService interface {
	Execute(purchase CreatePurchaseRequest) (Purchase, error)
}
