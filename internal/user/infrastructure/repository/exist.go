package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/user/domain"
)

func (repo UserMongoDBRepository) FindById(ID uuid.UUID) (*domain.User, error) {
	var userDb UserModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&userDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("product with id" + ID.String() + " does not exist")
		}
		log.Print(err)
		return nil, err
	}
	return MapToUserDomain(&userDb)
}
