package user

import (
	"errors"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain"
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

func MapToUserDomain(userModel *UserModel) (*domain.User, error) {
	user, err := domain.NewUser(userModel.ID, userModel.Name, userModel.Email, userModel.Surname)
	if err != nil {
		return nil, errors.New("cannot map user from db to model")
	}
	return user, nil
}
