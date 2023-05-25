package query

import (
	"meliarqsoft2/internal/domain/model"
)

type FindSellerEvent struct {
	repository model.ISellerRepository
}

func NewFindSellerEvent(repository model.ISellerRepository) *FindSellerEvent {
	return &FindSellerEvent{
		repository: repository,
	}
}

func (queryEvent FindSellerEvent) Execute(businessName string) ([]*model.Seller, error) {
	return queryEvent.repository.Find(businessName)
}
