package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"meliarqsoft2/internal/user/domain"
)

func (repo UserMongoDBRepository) Find(emailPattern string) ([]*domain.User, error) {
	var filter = bson.D{{
		Key: "email", Value: primitive.Regex{
			Pattern: ".*" + emailPattern + ".*",
			Options: "i",
		}}}

	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var dbResult []*UserModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		log.Print(err)
		return nil, err
	}

	var res []*domain.User
	for _, elem := range dbResult {
		res = append(res, MapToUserDomain(elem))
	}

	return res, nil
}
