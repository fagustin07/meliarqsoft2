package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreatePurchaseRequest struct {
	IDProduct uuid.UUID `json:"id_product" bson:"id_product,omitempty" binding:"required"`
	IDUser    uuid.UUID `json:"id_user" bson:"id_user,omitempty" binding:"required"`
	Units     int       `json:"units" bson:"units" binding:"required,gt=0"`
}

type FindPurchaseRequest struct {
	Start time.Time `json:"start" bson:"start"`
	Limit time.Time `json:"limit" bson:"limit"`
}
