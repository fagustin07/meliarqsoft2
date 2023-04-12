package application

import "github.com/google/uuid"

func (service PurchaseService) DeleteMany(productId uuid.UUID) error {
	return service.repository.DeleteMany(productId)
}
