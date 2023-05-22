package seller

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (repo MongoRepository) Restore(ID uuid.UUID) error {
	filter := bson.M{"_id": ID}

	updater := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "deleted_at",
					Value: time.Time{},
				},
			},
		},
	}

	_, err := repo.collection.UpdateOne(context.Background(), filter, updater)

	if err != nil {
		return err
	}

	return nil
}
