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
	"meliarqsoft2/internal/infrastructure/repository/mongo/purchase"
	"meliarqsoft2/internal/infrastructure/repository/mongo/seller"
	"meliarqsoft2/internal/infrastructure/repository/mongo/user"
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

	sellerRepository := seller.NewMongoRepository(client)
	userRepository := user.NewMongoRepository(client)
	productRepository := product.NewMongoRepository(client)
	purchaseRepository := purchase.NewPurchaseMongoRepository(client)

	// needed to build seller services
	findPurchasesFromProductEvent := query.NewFindPurchasesFromProductEvent(purchaseRepository)
	undoPurchase := action.NewUndoPurchaseCommand(purchaseRepository)
	undoPurchaseByProduct := action.NewUndoPurchasesByProductCommand(findPurchasesFromProductEvent, undoPurchase)
	findProduct := query.NewFindProductsBySellerCommand(productRepository)
	deleteProductEvent := action.NewDeleteProductEvent(productRepository, undoPurchaseByProduct)

	// seller
	deleteProductsBySeller := action.NewDeleteProductsBySellerCommand(productRepository, findProduct, deleteProductEvent)
	existSeller := query.NewExistSellerCommand(sellerRepository)
	registerSellerEvent := action.NewRegisterSellerEvent(sellerRepository)
	updateSellerEvent := action.NewUpdateSellerEvent(sellerRepository)
	unregisterSellerEvent := action.NewUnregisterSellerEvent(sellerRepository, deleteProductsBySeller, existSeller)
	findSellerEvent := query.NewFindSellerEvent(sellerRepository)

	// user
	existUser := query.NewExistUserCommand(userRepository)
	registerUserEvent := action.NewRegisterUserEvent(userRepository)
	updateUserEvent := action.NewUpdateUserEvent(userRepository)
	findUserEvent := query.NewFindUserEvent(userRepository)
	unregisterUserEvent := action.NewUnregisterUserEvent(userRepository)

	// product
	createProductEvent := action.NewCreateProductEvent(productRepository, existSeller)
	updateProductEvent := action.NewUpdateProductEvent(productRepository)
	findProductEvent := query.NewFindProductEvent(productRepository)
	filterProdctEvent := query.NewFilterProductEvent(productRepository)
	// delete product event was declared previously

	// make a purchase
	findProductCommand := query.NewFindProductCommand(productRepository)
	manageProductStock := action.NewManageProductStockCommand(productRepository)
	makePurchaseEvent := action.NewMakePurchaseEvent(purchaseRepository, findProductCommand, existUser, manageProductStock)

	newAPI := gin.NewMeliAPI(
		port,
		api.NewEvents(
			registerSellerEvent,
			updateSellerEvent,
			unregisterSellerEvent,
			findSellerEvent,

			registerUserEvent,
			updateUserEvent,
			unregisterUserEvent,
			findUserEvent,

			createProductEvent,
			updateProductEvent,
			deleteProductEvent,
			findProductEvent,
			filterProdctEvent,

			makePurchaseEvent,
			findPurchasesFromProductEvent,
		),
	)

	err = newAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}
