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
	"meliarqsoft2/internal/infrastructure/service"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port, err := strconv.Atoi(os.Getenv("PRODUCTS_APP_PORT"))
	if err != nil {
		panic(err.Error())
	}

	newFactory := mongo.NewFactory()
	client := newFactory.InitMongoDB()
	productRepository := product.NewMongoRepository(client)

	findSeller := service.NewSyncHttpFindSellerById(os.Getenv("USER_URL"))
	createProductEvent := action.NewCreateProductEvent(productRepository, findSeller)
	updateProductEvent := action.NewUpdateProductEvent(productRepository)
	findProductEvent := query.NewFindProductEvent(productRepository)
	filterProductEvent := query.NewFilterProductEvent(productRepository)
	findById := query.NewFindProductByIdEvent(productRepository)

	deletePurchases := service.NewSyncHttpDeletePurchasesByProducts(os.Getenv("PURCHASE_URL"))
	deleteProductEvent := action.NewDeleteProductEvent(productRepository, deletePurchases)
	deleteProductsBySeller := action.NewDeleteProductsBySellerEvent(productRepository, deletePurchases)

	newAPI := gin.NewMeliAPI(
		port,
		api.NewEvents(createProductEvent, updateProductEvent, deleteProductEvent, findProductEvent, filterProductEvent, deleteProductsBySeller, findById),
	)

	err = newAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}
