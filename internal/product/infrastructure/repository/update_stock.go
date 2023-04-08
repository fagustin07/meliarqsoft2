package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo ProductMongoDBRepository) UpdateStock(ID uuid.UUID, stock int) error {
	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{{"$set", bson.E{Key: "stock", Value: stock}}},
	)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
