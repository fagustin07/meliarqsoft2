package repository

import (
	"context"
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
		res = append(res, MapToPurchaseDomain(elem))
	}

	return res, err
}
