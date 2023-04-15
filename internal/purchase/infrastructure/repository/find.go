package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/purchase/domain"
)

func (repository PurchaseMongoDBRepository) Find(productID uuid.UUID) ([]*domain.Purchase, error) {
	filter := bson.D{{"id_product", productID}}

	cursor, err := repository.collection.Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Print(err)
	}

	var dbResult []*PurchaseModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		log.Print(err)
		return nil, err
	}

	var res []*domain.Purchase
	for _, elem := range dbResult {
		log.Print(elem)
		purchase, err := MapToPurchaseDomain(elem)
		if err != nil {
			return nil, errors.New("failed mapping purchase from db to model")
		}
		res = append(res, purchase)
	}

	return res, err
}
