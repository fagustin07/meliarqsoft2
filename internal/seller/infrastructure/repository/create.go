package repository

import (
	"context"
	"log"
	"meliarqsoft2/internal/seller/domain"
)

func (repo SellerMongoDBRepository) Create(seller *domain.Seller) error {
	dbSeller := mapSellerToMongoModel(seller)
	_, err := repo.collection.InsertOne(context.Background(), dbSeller)
	if err != nil {
		log.Println("Error saving seller", err)
		return err
	}

	return nil
}

func mapSellerToMongoModel(seller *domain.Seller) *SellerModel {
	return &SellerModel{ID: seller.ID, BusinessName: seller.BusinessName, Email: seller.BusinessName}
}
