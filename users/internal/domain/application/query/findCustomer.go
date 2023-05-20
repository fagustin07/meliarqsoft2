package query

import (
	"meliarqsoft2/internal/domain/model"
)

type FindCustomerEvent struct {
	repository model.ICustomerRepository
}

func NewFindCustomerEvent(repository model.ICustomerRepository) *FindCustomerEvent {
	return &FindCustomerEvent{
		repository: repository,
	}
}
func (queryEvent FindCustomerEvent) Execute(emailPattern string) ([]*model.Customer, error) {
	return queryEvent.repository.Find(emailPattern)
}
