package seller

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
)

type SellerModel struct {
	ID           uuid.UUID `json:"id" bson:"_id,omitempty"`
	BusinessName string    `json:"business_name" bson:"business_name"`
	Email        string    `json:"email" bson:"email"`
}

func MapSellerToMongoModel(seller *domain.Seller) *SellerModel {
	return &SellerModel{ID: seller.ID, BusinessName: seller.BusinessName, Email: seller.Email.Address}
}

func MapSellerFromModel(model *SellerModel) (*domain.Seller, error) {
	return domain.NewSeller(model.ID, model.BusinessName, model.Email)
}
