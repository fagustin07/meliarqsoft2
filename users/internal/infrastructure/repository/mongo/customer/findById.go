package customer

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meliarqsoft2/internal/domain/model"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) FindById(ID uuid.UUID) (*model.Customer, error) {
	var userDb Model
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&userDb)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model2.CustomerNotFoundError{Id: ID.String()}
		}
		return nil, err
	}
	return MapToCustomerDomain(&userDb)
}
