package dto

import "meliarqsoft2/internal/domain/model"

type CreateCustomerRequest struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Surname string `json:"surname" bson:"surname" binding:"required"`
	Email   string `json:"email" bson:"email" binding:"required,email"`
}

func (dto CreateCustomerRequest) MapToModel() (model.Customer, error) {
	address, err := model.NewEmail(dto.Email)
	if err != nil {
		return model.Customer{}, err
	}

	return model.Customer{
		Name:    dto.Name,
		Surname: dto.Surname,
		Email:   address,
	}, nil
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Surname string `json:"surname" bson:"surname" binding:"required"`
	Email   string `json:"email" bson:"email" binding:"required,email"`
}
