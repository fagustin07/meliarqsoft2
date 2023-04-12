package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (repo ProductMongoDBRepository) DeleteMany(sellerId uuid.UUID) error {
	var filter = bson.M{"id_seller": sellerId}

	_, err := repo.collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Print(err)
	}

	return err
}
