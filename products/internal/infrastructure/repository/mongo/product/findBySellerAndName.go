package product

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) FindBySellerAndName(sellerID uuid.UUID, name string) (*model.Product, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.M{"seller_id": sellerID},
				bson.M{"name": name},
			},
		},
	}

	var product model.Product
	if err := repo.collection.FindOne(context.Background(), filter).Decode(&product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model2.ProductNotFound{Id: ""}
		}
		return nil, err
	}
	return &product, nil
}
