package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo SellerMongoDBRepository) Exist(idSeller uuid.UUID) error {
	var sellerDb SellerModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": idSeller}).Decode(&sellerDb)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
