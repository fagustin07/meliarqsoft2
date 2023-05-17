package seller

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Create(seller *model.Seller) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, model2.CreateUUIDError{}
	}
	seller.ID = newUUID

	dbSeller := MapSellerToMongoModel(seller)
	_, err = repo.collection.InsertOne(context.Background(), dbSeller)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return uuid.Nil, model2.SellerAlreadyExist{}
		}

		return uuid.Nil, err
	}

	return newUUID, nil
}
