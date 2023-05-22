package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func (repo MongoRepository) DeleteBySeller(sellerID uuid.UUID) ([]uuid.UUID, error) {
	IDs, err := repo.FindIdsBySellerId(sellerID)
	if err != nil {
		return nil, err
	}

	log.Println(IDs)
	if len(IDs) == 0 {
		return nil, nil
	}

	filter := bson.M{"_id": bson.M{"$in": IDs}}

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

	_, err = repo.collection.UpdateMany(context.Background(), filter, updater)

	if err != nil {
		return nil, err
	}

	return IDs, nil
}
