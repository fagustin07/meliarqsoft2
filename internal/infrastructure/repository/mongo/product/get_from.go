package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"meliarqsoft2/internal/domain"
)

func (repo MongoRepository) GetFrom(sellerId uuid.UUID) ([]*domain.Product, error) {
	var filter = bson.M{"id_seller": sellerId}

	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var dbResult []*ProductModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		log.Print(err)
		return nil, err
	}

	var res []*domain.Product
	for _, elem := range dbResult {
		prod, err := mapProductToDomainModel(elem)
		if err != nil {
			return nil, err
		}
		res = append(res, prod)
	}

	return res, err
}
