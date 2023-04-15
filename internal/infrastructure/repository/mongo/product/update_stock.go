package product

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo MongoRepository) UpdateStock(ID uuid.UUID, stock int) error {
	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{{"$set", bson.D{bson.E{Key: "stock", Value: stock}}}},
	)
	if err != nil {
		log.Print(err)
		return errors.New("failed to update product stock")
	}

	return nil
}
