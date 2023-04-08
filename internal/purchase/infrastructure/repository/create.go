package repository

import (
	"context"
	"meliarqsoft2/internal/purchase/domain"
)

func (repository PurchaseMongoDBRepository) Create(purchase *domain.Purchase) error {
	dbPurchase := MapPurchaseToMongoModel(purchase)

	_, err := repository.collection.InsertOne(context.Background(), dbPurchase)

	if err != nil {
		return err
	}

	return nil
}
