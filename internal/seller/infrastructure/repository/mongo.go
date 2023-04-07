package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type SellerMongoDBRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewSellerMongoDBRepository(client *mongo.Client) *SellerMongoDBRepository {
	DbName := os.Getenv("DB_NAME")
	const DbCollection = "sellers"

	return &SellerMongoDBRepository{
		client:     client,
		database:   client.Database(DbName),
		collection: client.Database(DbName).Collection(DbCollection),
	}
}
