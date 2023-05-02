package user

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Delete(ID uuid.UUID) error {
	_, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": ID})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.UserNotFoundError{Id: ID.String()}
		}
		return err
	}

	return nil
}
