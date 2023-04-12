package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"meliarqsoft2/internal/product/domain"
)

func (repo ProductMongoDBRepository) GetProduct(ID uuid.UUID) (*domain.Product, error) {
	var productDb *ProductModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&productDb)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapProductToDomainModel(productDb)
}
