package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo UserMongoDBRepository) Exist(idUser uuid.UUID) error {
	var userDb UserModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": idUser}).Decode(&userDb)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
