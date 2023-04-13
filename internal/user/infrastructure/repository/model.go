package repository

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/user/domain"
)

type UserModel struct {
	ID      uuid.UUID `json:"id" bson:"_id,omitempty"`
	Name    string    `json:"name" bson:"name"`
	Surname string    `json:"surname" bson:"surname"`
	Email   string    `json:"email" bson:"email"`
}

func MapUserToMongoModel(user *domain.User) UserModel {
	return UserModel{ID: user.ID, Name: user.Name, Surname: user.Surname, Email: user.Email.Address}
}

func MapToUserDomain(userModel *UserModel) *domain.User {
	user, err := domain.NewUser(userModel.ID, userModel.Name, userModel.Email, userModel.Surname)
	if err != nil {
		return nil
	}
	return user
}
