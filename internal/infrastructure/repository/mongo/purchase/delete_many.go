package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repository MongoRepository) DeleteMany(ID uuid.UUID) error {
	var filter = bson.M{"product_id": ID}

	_, err := repository.collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Print(err)
	}

	return err
}
