package purchase

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/pkg/exceptions/model"
)

func (repository MongoRepository) Delete(ID uuid.UUID) error {
	_, err := repository.collection.DeleteOne(context.Background(), bson.M{"_id": ID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.PurchaseNotFoundError{Id: ID.String()}
		}
		return err
	}

	return nil
}
