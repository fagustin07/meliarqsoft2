package repository

import (
	"github.com/google/uuid"
)

type SellerModel struct {
	ID           uuid.UUID `json:"id" bson:"_id,omitempty"`
	BusinessName string    `json:"business_name" bson:"business_name"`
	Email        string    `json:"email" bson:"email"`
}
