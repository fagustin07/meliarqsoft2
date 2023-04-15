package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (repo SellerMongoDBRepository) Exist(idSeller uuid.UUID) (bool, error) {
	var sellerDb SellerModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": idSeller}).Decode(&sellerDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Print(err)
		return false, err
	}

	return true, nil
}
