package main

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"log"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/domain/model"
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

	if os.Getenv("ENVIRONMENT") == "test" {
		seed(productRepository)
	}

	err = newAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}

func seed(repository model.IProductRepository) {
	id, _ := uuid.Parse("9d032089-9223-4c98-9196-f97c6f792473")
	price, _ := model.NewPrice(float32(200))
	stock, _ := model.NewStock(999999)
	err := repository.DeleteAll()
	if err != nil && err != mongo2.ErrNoDocuments {
		panic("failed in product service seed")
	}

	_, err = repository.Create(model.Product{
		ID:          id,
		Name:        "prueba",
		Description: "prueba",
		Category:    "prueba",
		Price:       price,
		Stock:       stock,
		IDSeller:    id,
	})

	if err != nil {
		panic("failed in product service seed")
	}
}
