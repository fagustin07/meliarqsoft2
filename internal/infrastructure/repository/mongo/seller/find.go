package seller

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/domain/model"
)

func (repo MongoRepository) Find(businessName string) ([]*model.Seller, error) {
	var filter = bson.D{{
		Key: "business_name", Value: primitive.Regex{
			Pattern: ".*" + businessName + ".*",
			Options: "i",
		}}}

	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Print(err)
		return nil, nil
	}

	var dbResult []*SellerModel
	if err = cursor.All(context.TODO(), &dbResult); err != nil {
		log.Print(err)
		return nil, err
	}

	var res []*model.Seller
	for _, elem := range dbResult {
		res = append(res, mapProductToDomainModel(elem))
	}

	return res, nil
}

func mapProductToDomainModel(elem *SellerModel) *model.Seller {
	seller, err := model.NewSeller(elem.ID, elem.BusinessName, elem.Email)
	if err != nil {
		return nil
	}

	return seller
}
