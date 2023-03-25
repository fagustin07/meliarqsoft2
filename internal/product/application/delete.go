package application

import (
	"github.com/google/uuid"
)

func (service *ProductApplication) Delete(ID uuid.UUID) error {
	return service.repo.Delete(ID)
}
