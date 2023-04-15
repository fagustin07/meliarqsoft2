package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/seller/domain"
)

func (repo SellerMongoDBRepository) FindById(idSeller uuid.UUID) (*domain.Seller, error) {
	var sellerDb SellerModel
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": idSeller}).Decode(&sellerDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("seller not found")
		}
		log.Print(err)
		return nil, errors.New("failed while search seller")
	}

	return MapSellerFromModel(&sellerDb)
}
