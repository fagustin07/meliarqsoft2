package seller

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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
		log.Print(err)
		return err
	}

	return nil
}
