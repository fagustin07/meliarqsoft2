package user

import (
	"context"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Create(user *model.User) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, model2.CreateUUIDError{}
	}

	user.ID = newUUID
	_, err = repo.FindByEmail(user.Email.Address)
	if _, sellerNotFound := err.(model2.SellerNotFoundError); !sellerNotFound {
		return uuid.Nil, model2.UserAlreadyExistError{}
	}

	userDb := MapUserToMongoModel(user)
	_, err = repo.collection.InsertOne(context.Background(), userDb)
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil

}
