package model

import (
	"github.com/google/uuid"
)

type Seller struct {
	ID           uuid.UUID `json:"seller_id" bson:"seller_id"`
	BusinessName string    `json:"businessName" bson:"businessName"`
	Email        string    `json:"email" bson:"email"`
}

type CreateSellerRequest struct {
	BusinessName string `json:"business_name" bson:"business_name"`
	Email        string `json:"email" bson:"email"`
}

type SellerID struct {
	ID uuid.UUID `json:"seller_id" bson:"seller_id"`
}

type UpdateSellerRequest struct {
	BusinessName string `json:"business_name" bson:"business_name"`
	Email        string `json:"email" bson:"email"`
}

type ISellerService interface {
	Register(seller CreateSellerRequest) (uuid.UUID, error)
	Update(id uuid.UUID, businessName string, email string) error
	Find(businessName string) ([]Seller, error)
	Delete(ID uuid.UUID) error
}
