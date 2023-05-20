package customer

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Delete(ID uuid.UUID) error {
	customer, err := repo.FindById(ID)
	if err != nil {
		return err
	}

	if customer == nil {
		return model.CustomerNotFoundError{}
	}

	res, err := repo.collection.DeleteOne(context.Background(), bson.M{"_id": ID})

	if res.DeletedCount == 0 {
		return errors.New("cannot delete customer " + ID.String())
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.CustomerNotFoundError{Id: ID.String()}
		}
		return err
	}

	return nil
}
