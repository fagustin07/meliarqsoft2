package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
	"time"
)

func (repo MongoRepository) FindById(ID uuid.UUID) (*model.Product, error) {
	filter := bson.M{
		"_id":        ID,
		"deleted_at": time.Time{},
	}

	var productDb *ProductModel
	err := repo.collection.FindOne(context.Background(), filter).Decode(&productDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model2.ProductNotFound{Id: ID.String()}
		}
		return nil, err
	}

	product, err := mapProductToDomainModel(productDb)
	if err != nil {
		return nil, err
	}

	return &product, err
}
