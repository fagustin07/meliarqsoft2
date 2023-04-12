package application

import "github.com/google/uuid"

func (service SellerService) Update(id uuid.UUID, businessName string, email string) error {
	return service.repo.Update(id, businessName, email)
}
