package query

import (
	"meliarqsoft2/internal/seller/domain"
	"meliarqsoft2/internal/seller/domain/ports"
)

type FindSeller struct {
	repository ports.ISellerRepository
}

func (eventQuery FindSeller) Execute(businessName string) ([]*domain.Seller, error) {
	return eventQuery.repository.Find(businessName)
}
