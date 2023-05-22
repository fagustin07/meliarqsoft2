package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"time"
)

func (repo MongoRepository) UpdateStock(ID uuid.UUID, stock int) error {
	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID, "deleted_at": time.Time{}},
		bson.D{{"$set", bson.D{bson.E{Key: "stock", Value: stock}}}},
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model2.ProductNotFound{Id: ID.String()}
		}
		return err
	}

	return nil
}
