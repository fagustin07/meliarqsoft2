package application

import "github.com/google/uuid"

func (manager PurchaseService) DeleteMany(productId uuid.UUID) error {
	return manager.repository.DeleteMany(productId)
}
