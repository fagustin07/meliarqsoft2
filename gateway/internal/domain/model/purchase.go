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
	IDProduct uuid.UUID `json:"id_product" bson:"id_product,omitempty"`
	IDUser    uuid.UUID `json:"id_user" bson:"id_user,omitempty"`
	Units     int       `json:"units" bson:"units"`
	Total     float32   `json:"total" bson:"total"`
}

type IFindPurchaseService interface {
	Execute(productID uuid.UUID) ([]Purchase, error)
}
