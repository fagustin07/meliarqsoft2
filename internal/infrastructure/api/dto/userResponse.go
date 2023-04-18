package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
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

func MapUserToJSON(user *domain.User) UserDTO {
	return UserDTO{user.ID, user.Name, user.Surname, user.Email.Address}
}
