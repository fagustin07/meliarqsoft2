package user

import (
	"context"
	"github.com/google/uuid"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) Create(user *model.User) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	user.ID = newUUID

	userDb := MapUserToMongoModel(user)
	_, err = repo.collection.InsertOne(context.Background(), userDb)
	if err != nil {
		log.Println("Error saving user", err)
		return uuid.Nil, err
	}

	return newUUID, nil

}
