package user

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type UserModel struct {
	ID      uuid.UUID `json:"id" bson:"_id,omitempty"`
	Name    string    `json:"name" bson:"name"`
	Surname string    `json:"surname" bson:"surname"`
	Email   string    `json:"email" bson:"email"`
}

func MapUserToMongoModel(user *model.User) UserModel {
	return UserModel{ID: user.ID, Name: user.Name, Surname: user.Surname, Email: user.Email.Address}
}

func MapToUserDomain(userModel *UserModel) (*model.User, error) {
	address, err := model.NewEmail(userModel.Email)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:      userModel.ID,
		Name:    userModel.Name,
		Surname: userModel.Surname,
		Email:   address,
	}, nil
}
