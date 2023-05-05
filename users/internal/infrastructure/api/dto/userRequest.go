package dto

import "meliarqsoft2/internal/domain/model"

type CreateUserRequest struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Surname string `json:"surname" bson:"surname" binding:"required"`
	Email   string `json:"email" bson:"email" binding:"required,email"`
}

func (dto CreateUserRequest) MapToModel() (model.User, error) {
	address, err := model.NewEmail(dto.Email)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Name:    dto.Name,
		Surname: dto.Surname,
		Email:   address,
	}, nil
}

type UpdateUserRequest struct {
	Name    string `json:"name" bson:"name" binding:"required"`
	Surname string `json:"surname" bson:"surname" binding:"required"`
	Email   string `json:"email" bson:"email" binding:"required,email"`
}
