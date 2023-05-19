package customer

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type Model struct {
	ID      uuid.UUID `json:"id" bson:"_id,omitempty"`
	Name    string    `json:"name" bson:"name"`
	Surname string    `json:"surname" bson:"surname"`
	Email   string    `json:"email" bson:"email"`
}

func MapCustomerToMongoModel(user *model.Customer) Model {
	return Model{ID: user.ID, Name: user.Name, Surname: user.Surname, Email: user.Email.Address}
}

func MapToCustomerDomain(customerModel *Model) (*model.Customer, error) {
	address, err := model.NewEmail(customerModel.Email)
	if err != nil {
		return nil, err
	}

	return &model.Customer{
		ID:      customerModel.ID,
		Name:    customerModel.Name,
		Surname: customerModel.Surname,
		Email:   address,
	}, nil
}
