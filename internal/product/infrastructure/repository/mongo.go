package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type ProductMongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewProductMongoDB(client *mongo.Client) *ProductMongoDB {
	DbName := os.Getenv("DB_NAME")
	const DbCollection = "products"

	return &ProductMongoDB{
		client:     client,
		database:   client.Database(DbName),
		collection: client.Database(DbName).Collection(DbCollection),
	}
}
