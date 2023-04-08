package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/purchase/domain"
	"time"
)

func (repository PurchaseMongoDBRepository) Create(purchase *domain.Purchase) error {
	dbPurchase := MapPurchaseToMongoModel(purchase)

	_, err := repository.collection.InsertOne(context.Background(), dbPurchase)

	if err != nil {
		return err
	}

	return nil
}

func (repository PurchaseMongoDBRepository) Find(start time.Time, limit time.Time) ([]*domain.Purchase, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"date", bson.D{{"$gte", start}}}},
				bson.D{{"date", bson.D{{"$lte", limit}}}},
			},
		},
	}

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
		res = append(res, MapToPurchaseDomain(elem))
	}

	return res, err
}
