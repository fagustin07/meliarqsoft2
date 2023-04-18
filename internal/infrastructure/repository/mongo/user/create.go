package user

import (
	"context"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) Create(user *model.User) error {
	userDb := MapUserToMongoModel(user)
	_, err := repo.collection.InsertOne(context.Background(), userDb)
	if err != nil {
		log.Println("Error saving user", err)
		return err
	}

	return nil

}
