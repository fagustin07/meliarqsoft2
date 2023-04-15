package query

import (
	"meliarqsoft2/internal/domain"
)

type FindSellerEvent struct {
	repository domain.ISellerRepository
}

func NewFindSellerEvent(repository domain.ISellerRepository) *FindSellerEvent {
	return &FindSellerEvent{
		repository: repository,
	}
}

func (queryEvent FindSellerEvent) Execute(businessName string) ([]*domain.Seller, error) {
	return queryEvent.repository.Find(businessName)
}
