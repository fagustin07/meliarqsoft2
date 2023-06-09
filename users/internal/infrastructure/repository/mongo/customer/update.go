package customer

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Update(ID uuid.UUID, name string, surname string, email string) error {
	user, err := repo.FindById(ID)
	if err != nil {
		return err
	}

	if user == nil {
		return model2.CustomerNotFoundError{}
	}

	var fieldsToUpdate bson.D
	if name != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "name", Value: name})
	}

	if surname != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "surname", Value: surname})
	}

	if email != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "email", Value: email})
	}

	_, err = repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{{"$set", fieldsToUpdate}},
	)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model2.CustomerNotFoundError{Id: ID.String()}
		}
		if mongo.IsDuplicateKeyError(err) {
			return model2.CustomerAlreadyExistError{}
		}
		return err
	}

	return nil
}
