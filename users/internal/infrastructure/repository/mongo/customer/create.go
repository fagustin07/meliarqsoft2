package customer

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Create(customer *model.Customer) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, model2.CreateUUIDError{}
	}

	customer.ID = newUUID

	customerDb := MapCustomerToMongoModel(customer)
	_, err = repo.collection.InsertOne(context.Background(), customerDb)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return uuid.Nil, model2.CustomerAlreadyExistError{}
		}
		return uuid.Nil, err
	}

	return newUUID, nil

}
