package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo UserMongoDBRepository) Update(ID uuid.UUID, name string, surname string, email string) error {
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

	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{{"$set", fieldsToUpdate}},
	)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
