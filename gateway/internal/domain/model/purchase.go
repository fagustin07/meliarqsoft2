package model

import (
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	ID        uuid.UUID `json:"id" bson:"id,omitempty"`
	IDProduct uuid.UUID `json:"id_product" bson:"id_product,omitempty"`
	IDUser    uuid.UUID `json:"id_user" bson:"id_user,omitempty"`
	Date      time.Time `json:"date" bson:"date"`
	Units     int       `json:"units" bson:"units"`
	Total     float32   `json:"total" bson:"total"`
}

type CreatePurchaseRequest struct {
	IDProduct uuid.UUID `json:"product_id" bson:"product_id,omitempty"`
	IDUser    uuid.UUID `json:"customer_id" bson:"customer_id,omitempty"`
	Units     int       `json:"units" bson:"units"`
}

type IFindPurchaseService interface {
	Execute(productID uuid.UUID) ([]Purchase, error)
}

type IMakePurchaseService interface {
	Execute(purchase CreatePurchaseRequest) (Purchase, error)
}
