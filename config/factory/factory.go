package factory

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"meliarqsoft2/internal/product/application"
	"meliarqsoft2/internal/product/infrastructure/handler"
	"meliarqsoft2/internal/product/infrastructure/repository"
	"os"
	"time"
)

const MongoClientTimeout = 10

type Factory struct {
	dbClient *mongo.Client
}

func NewFactory() *Factory {
	return &Factory{}
}

func (factory *Factory) InitMongoDB() *mongo.Client {
	if factory.dbClient != nil {
		return factory.dbClient
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	factory.dbClient = client
	return client
}

func (factory *Factory) BuildProductHandler() *handler.ProductGinHandler {
	repo := repository.NewProductMongoDBRepository(factory.dbClient)
	app := application.NewProductApplication(repo)
	return handler.NewProductGinHandler(app)
}
