package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type ProductMongoDBRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewProductMongoDBRepository(client *mongo.Client) *ProductMongoDBRepository {
	DbName := os.Getenv("DB_NAME")
	const DbCollection = "products"

	return &ProductMongoDBRepository{
		client:     client,
		database:   client.Database(DbName),
		collection: client.Database(DbName).Collection(DbCollection),
	}
}
