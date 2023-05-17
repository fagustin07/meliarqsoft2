package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Delete(ID uuid.UUID) error {
	user, err := repo.FindById(ID)
	if err != nil {
		return err
	}

	if user == nil {
		return model.UserNotFoundError{}
	}

	res, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": ID})

	if res.DeletedCount == 0 {
		return errors.New("cannot delete user " + ID.String())
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.UserNotFoundError{Id: ID.String()}
		}
		return err
	}

	return nil
}
