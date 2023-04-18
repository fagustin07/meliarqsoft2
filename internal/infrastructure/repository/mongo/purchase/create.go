package purchase

import (
	"context"
	"meliarqsoft2/internal/domain/model"
)

func (repository MongoRepository) Create(purchase *model.Purchase) error {
	dbPurchase := MapPurchaseToMongoModel(purchase)

	_, err := repository.collection.InsertOne(context.Background(), dbPurchase)

	if err != nil {
		return err
	}

	return nil
}
