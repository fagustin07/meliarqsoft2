package seller

import (
	"context"
	"github.com/google/uuid"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Create(seller *model.Seller) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, model2.CreateUUIDError{}
	}
	seller.ID = newUUID

	_, err = repo.FindByEmail(seller.Email.Address)
	if _, sellerNotFound := err.(model2.SellerNotFoundError); !sellerNotFound {
		return uuid.Nil, model2.SellerAlreadyExist{}
	}

	_, err = repo.FindByBusinessName(seller.BusinessName)
	if _, sellerNotFound := err.(model2.SellerNotFoundError); !sellerNotFound {
		return uuid.Nil, model2.SellerAlreadyExist{}
	}

	dbSeller := MapSellerToMongoModel(seller)
	_, err = repo.collection.InsertOne(context.Background(), dbSeller)
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}
