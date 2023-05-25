package seller

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Create(seller *model.Seller) (uuid.UUID, error) {
	var newUUID uuid.UUID
	var err error

	if seller.ID == uuid.Nil {
		newUUID, err = uuid.NewUUID()
		if err != nil {
			return uuid.Nil, model2.CreateUUIDError{}
		}
	} else {
		newUUID = seller.ID
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

	log.Println("created seller with id " + newUUID.String())

	return newUUID, nil
}
