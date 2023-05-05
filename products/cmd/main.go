package main

import (
	"github.com/joho/godotenv"
	"log"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin"
	"meliarqsoft2/internal/infrastructure/repository/mongo"
	"meliarqsoft2/internal/infrastructure/repository/mongo/product"
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

	newFactory := mongo.NewFactory()
	client := newFactory.InitMongoDB()
	productRepository := product.NewMongoRepository(client)

	createProductEvent := action.NewCreateProductEvent(productRepository)
	updateProductEvent := action.NewUpdateProductEvent(productRepository)
	findProductEvent := query.NewFindProductEvent(productRepository)
	filterProductEvent := query.NewFilterProductEvent(productRepository)
	deleteProductEvent := action.NewDeleteProductEvent(productRepository)

	newAPI := gin.NewMeliAPI(
		port,
		api.NewEvents(
			createProductEvent,
			updateProductEvent,
			deleteProductEvent,
			findProductEvent,
			filterProductEvent,
		),
	)

	err = newAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}
