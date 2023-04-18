package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) FindById(ID uuid.UUID) (*model.User, error) {
	var userDb UserModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&userDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user with id " + ID.String() + " does not exist")
		}
		log.Print(err)
		return nil, err
	}
	return MapToUserDomain(&userDb)
}
