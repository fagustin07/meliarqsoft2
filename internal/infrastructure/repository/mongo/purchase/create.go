package purchase

import (
	"context"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
)

func (repository MongoRepository) Create(purchase *model.Purchase, product *model.Product) (uuid.UUID, float32, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, 0, err
	}
	newTotal, err := model.NewTotal(product.Price.Value * float32(purchase.Units.Amount))
	if err != nil {
		return uuid.Nil, 0, err
	}

	purchase.ID = newId
	purchase.Total = newTotal

	dbPurchase := MapPurchaseToMongoModel(purchase)
	_, err = repository.collection.InsertOne(context.Background(), dbPurchase)

	if err != nil {
		return uuid.Nil, 0, err
	}

	return newId, newTotal.Value, nil
}
