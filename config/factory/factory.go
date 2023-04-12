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
	application4 "meliarqsoft2/internal/purchase/application"
	handler4 "meliarqsoft2/internal/purchase/infrastructure/handler"
	repository4 "meliarqsoft2/internal/purchase/infrastructure/repository"
	application2 "meliarqsoft2/internal/seller/application"
	handler2 "meliarqsoft2/internal/seller/infrastructure/handler"
	repository2 "meliarqsoft2/internal/seller/infrastructure/repository"
	application3 "meliarqsoft2/internal/user/application"
	handler3 "meliarqsoft2/internal/user/infrastructure/handler"
	repository3 "meliarqsoft2/internal/user/infrastructure/repository"
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

func (factory *Factory) BuildProductHandler() (*application.ProductService, *handler.ProductGinHandler) {
	repo := repository.NewProductMongoDBRepository(factory.dbClient)
	purchaseService, _ := factory.BuildPurchaseHandler()
	sellerService, _ := factory.BuildSellerHandler()
	app := application.NewProductService(repo, sellerService, purchaseService)
	return app, handler.NewProductGinHandler(app)
}

func (factory *Factory) BuildSellerHandler() (*application2.SellerService, *handler2.SellerGinHandler) {
	repo := repository2.NewSellerMongoDBRepository(factory.dbClient)
	productService, _ := factory.BuildProductHandler()
	manager := application2.NewSellerService(repo, productService)
	return manager, handler2.NewSellerGinHandler(manager)
}

func (factory *Factory) BuildUserHandler() (*application3.UserService, *handler3.UserGinHandler) {
	repo := repository3.NewUserMongoDBRepository(factory.dbClient)
	manager := application3.NewUserService(repo)
	return manager, handler3.NewUserGinHandler(manager)
}

func (factory *Factory) BuildPurchaseHandler() (*application4.PurchaseService, *handler4.PurchaseGinHandler) {
	repo := repository4.NewPurchaseMongoDBRepository(factory.dbClient)
	prodService, _ := factory.BuildProductHandler()
	userService, _ := factory.BuildUserHandler()
	service := application4.NewPurchaseService(repo, prodService, userService)
	return service, handler4.NewPurchaseGinHandler(service)
}
