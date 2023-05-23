package query

import (
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

type FindCustomerByIdEvent struct {
	repository model.ICustomerRepository
}

func NewFindCustomerByIdEvent(repository model.ICustomerRepository) *FindCustomerByIdEvent {
	return &FindCustomerByIdEvent{
		repository: repository,
	}
}

func (queryEvent FindCustomerByIdEvent) Execute(id uuid.UUID) (*model.Customer, error) {
	return queryEvent.repository.FindById(id)
}
