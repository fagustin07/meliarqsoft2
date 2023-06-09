package product

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	"time"
)

func (repo MongoRepository) Find(name string, category string) ([]model.Product, error) {
	var filter = bson.D{
		{"$and",
			bson.A{
				bson.D{{
					Key: "name", Value: primitive.Regex{
						Pattern: ".*" + name + ".*",
						Options: "i",
					}}},
				bson.D{{
					Key: "category", Value: primitive.Regex{
						Pattern: ".*" + category + ".*",
						Options: "i",
					}}},
				bson.D{{
					Key: "deleted_at", Value: time.Time{}}},
			},
		},
	}
	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
	}

	var dbResult []*ProductModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		return nil, err
	}

	var res []model.Product
	for _, elem := range dbResult {
		prod, err := mapProductToDomainModel(elem)
		if err != nil {
			return nil, err
		}
		res = append(res, prod)
	}

	return res, err
}
