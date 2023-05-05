package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) FindById(ID uuid.UUID) (*model.Product, error) {
	var productDb *ProductModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&productDb)
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
