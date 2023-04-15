package user

import (
	"context"
	"log"
	"meliarqsoft2/internal/domain"
)

func (repo MongoRepository) Create(user *domain.User) error {
	userDb := MapUserToMongoModel(user)
	_, err := repo.collection.InsertOne(context.Background(), userDb)
	if err != nil {
		log.Println("Error saving user", err)
		return err
	}

	return nil

}
