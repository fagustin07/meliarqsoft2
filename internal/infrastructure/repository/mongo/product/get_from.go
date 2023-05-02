package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) GetFrom(sellerId uuid.UUID) ([]model.Product, error) {
	var filter = bson.M{"id_seller": sellerId}

	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var dbResult []*ProductModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		log.Print(err)
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
