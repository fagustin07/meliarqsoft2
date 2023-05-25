package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type SellerID struct {
	ID uuid.UUID `json:"seller_id" bson:"seller_id"`
}

type SellerDTO struct {
	ID           uuid.UUID `json:"seller_id" bson:"seller_id"`
	BusinessName string    `json:"businessName" bson:"businessName"`
	Email        string    `json:"email" bson:"email"`
}

func (dto SellerDTO) MapToModel() (model.Seller, error) {
	address, err := model.NewEmail(dto.Email)
	if err != nil {
		return model.Seller{}, err
	}

	return model.Seller{
		ID:           dto.ID,
		BusinessName: dto.BusinessName,
		Email:        address,
	}, nil
}

func MapSellerToJSON(seller *model.Seller) SellerDTO {
	return SellerDTO{seller.ID, seller.BusinessName, seller.Email.Address}
}
