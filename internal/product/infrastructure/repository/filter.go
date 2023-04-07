package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/product/domain"
)

func (repo ProductMongoDBRepository) Filter(minPrice float32, maxPrice float32) ([]*domain.Product, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"price", bson.D{{"$gte", minPrice}}}},
				bson.D{{"price", bson.D{{"$lte", maxPrice}}}},
			},
		},
	}

	cursor, err := repo.collection.Find(context.Background(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Print(err)
	}

	var dbResult []*ProductModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		log.Print(err)
		return nil, err
	}

	var res []*domain.Product
	for _, elem := range dbResult {
		res = append(res, mapProductToDomainModel(elem))
	}

	return res, err
}
