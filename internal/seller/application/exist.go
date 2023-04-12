package application

import "github.com/google/uuid"

func (service SellerService) Exist(seller uuid.UUID) error {
	return service.repo.Exist(seller)
}
