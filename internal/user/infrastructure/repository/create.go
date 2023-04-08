package repository

import (
	"context"
	"log"
	"meliarqsoft2/internal/user/domain"
)

func (repo UserMongoDBRepository) Create(user *domain.User) error {
	userDb := MapUserToMongoModel(user)
	_, err := repo.collection.InsertOne(context.Background(), userDb)
	if err != nil {
		log.Println("Error saving user", err)
		return err
	}

	return nil

}
