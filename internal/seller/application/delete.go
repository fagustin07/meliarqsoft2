package application

import (
	"github.com/google/uuid"
)

func (service SellerService) Delete(ID uuid.UUID) error {
	return service.repo.Delete(ID)
}
