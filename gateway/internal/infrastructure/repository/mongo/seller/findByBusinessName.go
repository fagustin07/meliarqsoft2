package seller

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) FindByBusinessName(businessName string) (*model.Seller, error) {
	filter := bson.M{"business_name": businessName}
	var seller model.Seller
	if err := repo.collection.FindOne(context.Background(), filter).Decode(&seller); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model2.SellerNotFoundError{Id: ""}
		}
		return nil, err
	}
	return &seller, nil
}
