package application

import (
	"github.com/google/uuid"
)

func (service SellerService) Delete(ID uuid.UUID) error {
	err := service.productService.DeleteMany(ID)
	if err != nil {
		return err
	}
	return service.repo.Delete(ID)
}
