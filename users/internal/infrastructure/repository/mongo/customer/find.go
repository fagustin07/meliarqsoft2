package customer

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) Find(emailPattern string) ([]*model.Customer, error) {
	var filter = bson.D{{
		Key: "email", Value: primitive.Regex{
			Pattern: ".*" + emailPattern + ".*",
			Options: "i",
		}}}

	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var dbResult []*Model
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		return nil, err
	}

	var res []*model.Customer
	for _, elem := range dbResult {
		user, err := MapToCustomerDomain(elem)
		if err != nil {
			return nil, err
		}

		res = append(res, user)
	}

	return res, nil
}
