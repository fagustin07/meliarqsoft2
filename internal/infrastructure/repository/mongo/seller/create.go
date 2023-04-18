package seller

import (
	"context"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) Create(seller *model.Seller) error {
	dbSeller := MapSellerToMongoModel(seller)
	_, err := repo.collection.InsertOne(context.Background(), dbSeller)
	if err != nil {
		log.Println("Error saving seller", err)
		return err
	}

	return nil
}
