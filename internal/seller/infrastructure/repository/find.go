package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/seller/domain"
)

func (repo SellerMongoDBRepository) Find(businessName string) ([]*domain.Seller, error) {
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

	var res []*domain.Seller
	for _, elem := range dbResult {
		res = append(res, mapProductToDomainModel(elem))
	}

	return res, nil
}

func mapProductToDomainModel(elem *SellerModel) *domain.Seller {
	return &domain.Seller{
		ID: elem.ID, BusinessName: elem.BusinessName, Email: elem.Email,
	}
}
