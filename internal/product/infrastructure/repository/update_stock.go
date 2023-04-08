package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo ProductMongoDBRepository) UpdateStock(ID uuid.UUID, stock int) error {
	x, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{{"$set", bson.D{bson.E{Key: "stock", Value: stock}}}},
	)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Println(x.UpsertedID)
	return nil
}
