package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (repo MongoRepository) Restore(IDs []uuid.UUID) (int64, error) {
	filter := bson.M{"_id": bson.M{"$in": IDs}}

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

	updRes, err := repo.collection.UpdateMany(context.Background(), filter, updater)

	if err != nil {
		return updRes.MatchedCount, err
	}

	return updRes.MatchedCount, nil
}
