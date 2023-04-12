package application

import "meliarqsoft2/internal/seller/domain"

func (service SellerService) Find(businessName string) ([]*domain.Seller, error) {
	return service.repo.Find(businessName)
}
