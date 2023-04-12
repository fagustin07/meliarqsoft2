package repository

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain"
)

type SellerModel struct {
	ID           uuid.UUID `json:"id" bson:"_id,omitempty"`
	BusinessName string    `json:"business_name" bson:"business_name"`
	Email        string    `json:"email" bson:"email"`
}

func mapSellerToMongoModel(seller *domain.Seller) *SellerModel {
	return &SellerModel{ID: seller.ID, BusinessName: seller.BusinessName, Email: seller.Email.Address}
}
