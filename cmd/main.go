package main

import (
	"github.com/joho/godotenv"
	"log"
	action2 "meliarqsoft2/internal/domain/application/command/action"
	query2 "meliarqsoft2/internal/domain/application/command/query"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin"
	"meliarqsoft2/internal/infrastructure/repository/mongo"
	"meliarqsoft2/internal/infrastructure/repository/mongo/product"
	"meliarqsoft2/internal/infrastructure/repository/mongo/purchase"
	"meliarqsoft2/internal/infrastructure/repository/mongo/seller"
	"meliarqsoft2/internal/infrastructure/repository/mongo/user"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newFactory := mongo.NewFactory()

	client := newFactory.InitMongoDB()

	sellerRepository := seller.NewMongoRepository(client)
	userRepository := user.NewMongoRepository(client)
	productRepository := product.NewMongoRepository(client)
	purchaseRepository := purchase.NewPurchaseMongoRepository(client)

	// needed to build seller services
	findPurchasesFromProductEvent := query2.NewFindPurchasesFromProductEvent(purchaseRepository)
	undoPurchase := action2.NewUndoPurchaseCommand(purchaseRepository)
	undoPurchaseByProduct := action2.NewUndoPurchasesByProductCommand(findPurchasesFromProductEvent, undoPurchase)
	findProduct := query2.NewFindProductsBySellerCommand(productRepository)
	deleteProductEvent := action2.NewDeleteProductEvent(productRepository, undoPurchaseByProduct)

	// seller
	deleteProductsBySeller := action2.NewDeleteProductsBySellerCommand(productRepository, findProduct, deleteProductEvent)
	existSeller := query2.NewExistSellerCommand(sellerRepository)
	registerSellerEvent := action2.NewRegisterSellerEvent(sellerRepository)
	updateSellerEvent := action2.NewUpdateSellerEvent(sellerRepository)
	unregisterSellerEvent := action2.NewUnregisterSellerEvent(sellerRepository, deleteProductsBySeller, existSeller)
	findSellerEvent := query2.NewFindSellerEvent(sellerRepository)

	// user
	existUser := query2.NewExistUserCommand(userRepository)
	registerUserEvent := action2.NewRegisterUserEvent(userRepository)
	updateUserEvent := action2.NewUpdateUserEvent(userRepository)
	findUserEvent := query2.NewFindUserEvent(userRepository)
	unregisterUserEvent := action2.NewUnregisterUserEvent(userRepository)

	// product
	createProductEvent := action2.NewCreateProductEvent(productRepository, existSeller)
	updateProductEvent := action2.NewUpdateProductEvent(productRepository)
	findProductEvent := query2.NewFindProductEvent(productRepository)
	filterProdctEvent := query2.NewFilterProductEvent(productRepository)
	// delete product event was declared previously

	// make a purchase
	findProductCommand := query2.NewFindProductCommand(productRepository)
	manageProductStock := action2.NewManageProductStockCommand(productRepository)
	makePurchaseEvent := action2.NewMakePurchaseEvent(purchaseRepository, findProductCommand, existUser, manageProductStock)

	newAPI := gin.NewMeliAPI(
		8080,
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
