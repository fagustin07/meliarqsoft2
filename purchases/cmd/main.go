package main

import (
	"github.com/joho/godotenv"
	"log"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin"
	"meliarqsoft2/internal/infrastructure/repository/mongo"
	"meliarqsoft2/internal/infrastructure/repository/mongo/purchase"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err.Error())
	}

	client := mongo.NewFactory().InitMongoDB()

	purchaseRepository := purchase.NewPurchaseMongoRepository(client)

	makePurchaseEvent := action.NewMakePurchaseEvent(purchaseRepository)

	undoPurchasesFromProductEvent := action.NewUndoPurchasesByProductEvent(
		query.NewFindPurchasesFromProductEvent(purchaseRepository),
		action.NewUndoPurchaseCommand(purchaseRepository),
	)

	findFromProductsEvent := query.NewFindPurchasesFromProductEvent(purchaseRepository)

	purchaseAPI := gin.NewMeliAPI(
		port,
		api.NewPurchaseEvents(
			undoPurchasesFromProductEvent,
			makePurchaseEvent,
			findFromProductsEvent,
		),
	)

	err = purchaseAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}
