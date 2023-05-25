package product

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/domain/model"
	"time"
)

func (repo MongoRepository) Filter(minPrice float32, maxPrice float32) ([]model.Product, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"price", bson.D{{"$gte", minPrice}}}},
				bson.D{{"price", bson.D{{"$lte", maxPrice}}}},
				bson.D{{"deleted_at", time.Time{}}},
			},
		},
	}

	cursor, err := repo.collection.Find(context.Background(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Print(err)
		return nil, err
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

	return res, nil
}
