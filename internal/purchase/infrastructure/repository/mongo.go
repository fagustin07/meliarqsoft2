package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type PurchaseMongoDBRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewPurchaseMongoDBRepository(client *mongo.Client) *PurchaseMongoDBRepository {
	DbName := os.Getenv("DB_NAME")
	const DbCollection = "purchases"

	return &PurchaseMongoDBRepository{
		client:     client,
		database:   client.Database(DbName),
		collection: client.Database(DbName).Collection(DbCollection),
	}
}
