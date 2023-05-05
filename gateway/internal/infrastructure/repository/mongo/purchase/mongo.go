package purchase

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type MongoRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewPurchaseMongoRepository(client *mongo.Client) *MongoRepository {
	DbName := os.Getenv("DB_NAME")
	const DbCollection = "purchases"

	return &MongoRepository{
		client:     client,
		database:   client.Database(DbName),
		collection: client.Database(DbName).Collection(DbCollection),
	}
}
