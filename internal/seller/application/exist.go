package application

import "github.com/google/uuid"

func (service SellerService) Exist(seller uuid.UUID) error {
	_, err := service.repo.FindById(seller)
	return err
}
