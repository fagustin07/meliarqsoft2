package user

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) Find(emailPattern string) ([]*model.User, error) {
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

	var res []*model.User
	for _, elem := range dbResult {
		user, err := MapToUserDomain(elem)
		if err != nil {
			return nil, errors.New("failed mapping user from db to model")
		}

		res = append(res, user)
	}

	return res, nil
}
