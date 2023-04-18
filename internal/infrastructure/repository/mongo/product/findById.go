package product

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) FindById(ID uuid.UUID) (*model.Product, error) {
	var productDb *ProductModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&productDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("product with id" + ID.String() + " does not exist")
		}
		log.Print(err)
		return nil, err
	}

	product, err := mapProductToDomainModel(productDb)
	if err != nil {
		return nil, err
	}

	return &product, err
}
