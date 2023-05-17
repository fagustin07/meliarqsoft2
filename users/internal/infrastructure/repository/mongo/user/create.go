package user

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Create(user *model.User) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, model2.CreateUUIDError{}
	}

	user.ID = newUUID

	userDb := MapUserToMongoModel(user)
	_, err = repo.collection.InsertOne(context.Background(), userDb)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return uuid.Nil, model2.UserAlreadyExistError{}
		}
		return uuid.Nil, err
	}

	return newUUID, nil

}
