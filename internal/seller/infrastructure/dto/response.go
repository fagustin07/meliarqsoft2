package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/seller/domain"
)

type SellerID struct {
	ID uuid.UUID `json:"seller_id" bson:"seller_id"`
}

type SellerDTO struct {
	ID           uuid.UUID `json:"seller_id" bson:"seller_id"`
	BusinessName string    `json:"businessName" bson:"businessName"`
	Email        string    `json:"email" bson:"email"`
}

func MapSellerToJSON(seller *domain.Seller) SellerDTO {
	return SellerDTO{seller.ID, seller.BusinessName, seller.Email.Address}
}
