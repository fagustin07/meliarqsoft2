package seller

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type SellerModel struct {
	ID           uuid.UUID `json:"id" bson:"_id,omitempty"`
	BusinessName string    `json:"business_name" bson:"business_name"`
	Email        string    `json:"email" bson:"email"`
}

func MapSellerToMongoModel(seller *model.Seller) *SellerModel {
	return &SellerModel{ID: seller.ID, BusinessName: seller.BusinessName, Email: seller.Email.Address}
}

func MapSellerFromModel(mongoModel *SellerModel) (*model.Seller, error) {
	address, err := model.NewEmail(mongoModel.Email)
	if err != nil {
		return nil, err
	}

	return &model.Seller{
		ID:           mongoModel.ID,
		BusinessName: mongoModel.BusinessName,
		Email:        address,
	}, nil
}
