package query

import (
	"meliarqsoft2/internal/seller/domain"
	"meliarqsoft2/internal/seller/domain/ports"
)

type FindSellerByBusinessNameEvent struct {
	repository ports.ISellerRepository
}

func (eventQuery FindSellerByBusinessNameEvent) Execute(businessName string) ([]*domain.Seller, error) {
	return eventQuery.repository.Find(businessName)
}
