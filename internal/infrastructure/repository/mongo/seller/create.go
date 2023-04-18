package seller

import (
	"context"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) Create(seller *model.Seller) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}
	seller.ID = newUUID

	dbSeller := MapSellerToMongoModel(seller)
	_, err = repo.collection.InsertOne(context.Background(), dbSeller)
	if err != nil {
		log.Println("Error saving seller", err)
		return uuid.Nil, err
	}

	return newUUID, nil
}
