package purchase

import (
	"context"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repository MongoRepository) Create(purchase *model.Purchase) (uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, model2.CreateUUIDError{}
	}

	purchase.ID = newId

	dbPurchase := MapPurchaseToMongoModel(purchase)
	_, err = repository.collection.InsertOne(context.Background(), dbPurchase)

	if err != nil {
		return uuid.Nil, err
	}

	return newId, nil
}
