package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"meliarqsoft2/internal/user/domain"
)

func (repo UserMongoDBRepository) Find(ID uuid.UUID) (*domain.User, error) {

	var userDb UserModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&userDb)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return MapToUserDomain(userDb), nil
}
