package seller

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	model2 "meliarqsoft2/pkg/exceptions/model"
)

func (repo MongoRepository) Update(ID uuid.UUID, businessName string, email string) error {
	var fieldsToUpdate bson.D
	if businessName != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "business_name", Value: businessName})
	}
	if email != "" {
		fieldsToUpdate = append(fieldsToUpdate, bson.E{Key: "email", Value: email})
	}

	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{{"$set", fieldsToUpdate}},
	)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model2.SellerNotFoundError{Id: ID.String()}
		}
		return err
	}

	return nil
}
