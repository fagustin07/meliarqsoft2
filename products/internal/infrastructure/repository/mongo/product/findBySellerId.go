package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func (repo MongoRepository) FindIdsBySellerId(sellerId uuid.UUID) ([]uuid.UUID, error) {
	log.Println(sellerId)
	sellerFilter := bson.M{
		"id_seller":  sellerId,
		"deleted_at": time.Time{},
	}

	cursor, err := repo.collection.Find(context.Background(), sellerFilter)

	var dbResult []*ProductModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		return nil, err
	}

	var res []uuid.UUID
	for _, elem := range dbResult {
		res = append(res, elem.ID)
	}

	return res, nil
}
