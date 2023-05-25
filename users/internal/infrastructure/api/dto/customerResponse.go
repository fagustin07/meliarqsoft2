package dto

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type CustomerID struct {
	ID uuid.UUID `json:"customer_id" bson:"customer_id"`
}

type CustomerDTO struct {
	ID      uuid.UUID `json:"customer_id" bson:"customer_id"`
	Name    string    `json:"name" bson:"name"`
	Surname string    `json:"surname" bson:"surname"`
	Email   string    `json:"email" bson:"email"`
}

func (dto CustomerDTO) MapToModel() (model.Customer, error) {
	address, err := model.NewEmail(dto.Email)
	if err != nil {
		return model.Customer{}, err
	}

	return model.Customer{
		ID:      dto.ID,
		Name:    dto.Name,
		Surname: dto.Surname,
		Email:   address,
	}, nil
}

func MapCustomerToJSON(customer *model.Customer) CustomerDTO {
	return CustomerDTO{customer.ID, customer.Name, customer.Surname, customer.Email.Address}
}
