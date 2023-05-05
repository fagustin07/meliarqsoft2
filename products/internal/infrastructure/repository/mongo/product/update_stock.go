package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo MongoRepository) UpdateStock(ID uuid.UUID, stock int) error {
	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{{"$set", bson.D{bson.E{Key: "stock", Value: stock}}}},
	)
	if err != nil {
		return err
	}

	return nil
}
