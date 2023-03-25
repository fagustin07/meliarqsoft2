package application

import (
	"github.com/google/uuid"
)

func (service *ProductService) Delete(ID uuid.UUID) error {
	return service.repo.Delete(ID)
}
