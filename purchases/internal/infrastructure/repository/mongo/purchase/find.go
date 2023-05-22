package purchase

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/domain/model"
	"time"
)

func (repository MongoRepository) Find(productID uuid.UUID) ([]*model.Purchase, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"id_product", productID}},
				bson.D{{"deleted_at", time.Time{}}},
			},
		},
	}
	cursor, err := repository.collection.Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Print(err)
		return nil, err
	}

	var dbResult []*PurchaseModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		log.Print(err)
		return nil, err
	}

	var res []*model.Purchase
	for _, elem := range dbResult {
		log.Print(elem)
		purchase, err := MapToPurchaseDomain(elem)
		if err != nil {
			return nil, err
		}
		res = append(res, purchase)
	}

	return res, nil
}
