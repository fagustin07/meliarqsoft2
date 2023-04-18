package dto

import "meliarqsoft2/internal/domain/model"

type CreateSellerRequest struct {
	BusinessName string `json:"business_name" bson:"business_name" binding:"required"`
	Email        string `json:"email" bson:"email" binding:"required,email"`
}

func (dto CreateSellerRequest) MapToModel() (model.Seller, error) {
	address, err := model.NewEmail(dto.Email)
	if err != nil {
		return model.Seller{}, err
	}

	return model.Seller{
		BusinessName: dto.BusinessName,
		Email:        address,
	}, nil
}

type UpdateSellerRequest struct {
	BusinessName string `json:"business_name" bson:"business_name" binding:"required"`
	Email        string `json:"email" bson:"email" binding:"required,email"`
}
