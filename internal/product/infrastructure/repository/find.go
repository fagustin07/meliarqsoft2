package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/product/domain"
)

func (repo ProductMongoDBRepository) Find(name string, category string) ([]*domain.Product, error) {
	var filter bson.M
	if name != "" {
		filter["name"] = "/" + name + "/"
	}

	if category != "" {
		filter["category"] = "/" + category + "/"
	}

	if filter == nil {
		return nil, nil
	} else {
		cursor, err := repo.collection.Find(context.Background(), filter)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, err
			}
			log.Fatal(err)
		}

		var dbResult []*ProductModel
		if err = cursor.All(context.TODO(), &dbResult); err != nil {
			log.Fatal(err)
			return nil, err
		}

		var res []*domain.Product
		for _, elem := range dbResult {
			res = append(res, mapProductToDomainModel(elem))
		}

		return res, err
	}
}
