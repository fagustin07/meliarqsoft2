package main

import (
	"github.com/joho/godotenv"
	"log"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin"
	"meliarqsoft2/internal/infrastructure/repository/mongo"
	"meliarqsoft2/internal/infrastructure/repository/mongo/purchase"
	"meliarqsoft2/internal/infrastructure/repository/rabbitmq"
	"meliarqsoft2/internal/infrastructure/repository/rabbitmq/notification"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port, err := strconv.Atoi(os.Getenv("PURCHASES_APP_PORT"))
	if err != nil {
		panic(err.Error())
	}

	client := mongo.NewFactory().InitMongoDB()

	purchaseRepository := purchase.NewPurchaseMongoRepository(client)

	// RabbitMQ
	clientMQ := rabbitmq.NewFactory().InitRabbitMQ()
	notificationService := notification.NewRabbitMQService(clientMQ)

	makePurchaseEvent := action.NewMakePurchaseEvent(purchaseRepository, notificationService)
	undoPurchasesFromProductEvent := action.NewUndoPurchasesByProductEvent(purchaseRepository)
	findFromProductsEvent := query.NewFindPurchasesFromProductEvent(purchaseRepository)

	purchaseAPI := gin.NewMeliAPI(
		port,
		api.NewPurchaseEvents(
			undoPurchasesFromProductEvent,
			makePurchaseEvent,
			findFromProductsEvent,
		),
	)

	if os.Getenv("ENVIRONMENT") == "test" {
		seed(purchaseRepository)
	}

	if err != nil {
		panic(err.Error())
	}

	err = purchaseAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}

func seed(repository model.IPurchaseRepository) {
	err := repository.DeleteAll()
	if err != nil {
		panic("error: " + err.Error())
	}
}
