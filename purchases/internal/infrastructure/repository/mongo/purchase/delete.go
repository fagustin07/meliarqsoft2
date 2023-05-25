package purchase

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (repository MongoRepository) DeleteByProductsIDs(productsIDs []uuid.UUID) error {
	filter := bson.M{"id_product": bson.M{"$in": productsIDs}}

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

	_, err := repository.collection.UpdateMany(context.Background(), filter, updater)

	if err != nil {
		return err
	}

	return nil
}
