package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"meliarqsoft2/internal/product/domain"
)

func (repo ProductMongoDBRepository) Update(ID uuid.UUID, name string, description string, category string, price float32, stock int) (*domain.Product, error) {
	_, err := repo.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": ID},
		bson.D{
			{"$set",
				bson.D{
					{"name", name},
					{"description", description},
					{"category", category},
					{"price", price},
					{"stock", stock},
				},
			},
		},
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return nil, nil
}
