package query

import (
	"meliarqsoft2/internal/seller/domain"
	"meliarqsoft2/internal/seller/domain/ports"
)

type FindSellerEvent struct {
	repository ports.ISellerRepository
}

func NewFindSellerEvent(repository ports.ISellerRepository) *FindSellerEvent {
	return &FindSellerEvent{
		repository: repository,
	}
}

func (queryEvent FindSellerEvent) Execute(businessName string) ([]*domain.Seller, error) {
	return queryEvent.repository.Find(businessName)
}
