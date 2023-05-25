package seller

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meliarqsoft2/internal/domain/model"
	"time"
)

func (repo MongoRepository) Find(businessName string) ([]*model.Seller, error) {
	var filter = bson.D{
		{"$and",
			bson.A{
				bson.D{{
					Key: "business_name", Value: primitive.Regex{
						Pattern: ".*" + businessName + ".*",
						Options: "i",
					}}},
				bson.D{{
					Key: "deleted_at", Value: time.Time{}}},
			},
		},
	}

	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var dbResult []*SellerModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		return nil, err
	}

	var res []*model.Seller
	for _, elem := range dbResult {
		seller, err2 := mapSellerToDomainModel(elem)
		if err2 != nil {
			return nil, err2
		}
		res = append(res, seller)
	}

	return res, nil
}

func mapSellerToDomainModel(elem *SellerModel) (*model.Seller, error) {
	address, err := model.NewEmail(elem.Email)
	if err != nil {
		return nil, err
	}

	return &model.Seller{
		ID:           elem.ID,
		BusinessName: elem.BusinessName,
		Email:        address,
	}, nil
}
