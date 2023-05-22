package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/pkg/exceptions/model"
	"time"
)

func (repo MongoRepository) Delete(ID uuid.UUID) error {
	filter := bson.M{
		"_id":        ID,
		"deleted_at": time.Time{},
	}

	updater := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "deleted_at",
					Value: time.Now(),
				},
			},
		},
	}

	_, err := repo.collection.UpdateOne(context.Background(), filter, updater)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.ProductNotFound{Id: ID.String()}
		}
		return err
	}

	return nil
}
