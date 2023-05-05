package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UserID struct {
	ID uuid.UUID `json:"user_id" bson:"seller_id"`
}

type UserDTO struct {
	ID      uuid.UUID `json:"user_id" bson:"user_id"`
	Name    string    `json:"name" bson:"name"`
	Surname string    `json:"surname" bson:"surname"`
	Email   string    `json:"email" bson:"email"`
}

func (dto UserDTO) MapToModel() (model.User, error) {
	address, err := model.NewEmail(dto.Email)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:      dto.ID,
		Name:    dto.Name,
		Surname: dto.Surname,
		Email:   address,
	}, nil
}

func MapUserToJSON(user *model.User) UserDTO {
	return UserDTO{user.ID, user.Name, user.Surname, user.Email.Address}
}