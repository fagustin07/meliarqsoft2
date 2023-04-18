package purchase

import (
	"context"
	"meliarqsoft2/internal/domain"
)

func (repository MongoRepository) Create(purchase *domain.Purchase) error {
	dbPurchase := MapPurchaseToMongoModel(purchase)

	_, err := repository.collection.InsertOne(context.Background(), dbPurchase)

	if err != nil {
		return err
	}

	return nil
}
