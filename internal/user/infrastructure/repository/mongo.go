package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type UserMongoDBRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserMongoDBRepository(client *mongo.Client) *UserMongoDBRepository {
	DbName := os.Getenv("DB_NAME")
	const DbCollection = "users"

	return &UserMongoDBRepository{
		client:     client,
		database:   client.Database(DbName),
		collection: client.Database(DbName).Collection(DbCollection),
	}
}
