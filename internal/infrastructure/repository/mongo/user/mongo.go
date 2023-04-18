package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	DbName := os.Getenv("DB_NAME")
	const DbCollection = "users"

	repo := &MongoRepository{
		client:     client,
		database:   client.Database(DbName),
		collection: client.Database(DbName).Collection(DbCollection),
	}

	_, err := repo.collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		panic("Failed creating user restrictions.")
	}

	return repo
}
