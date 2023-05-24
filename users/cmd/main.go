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
	"meliarqsoft2/internal/infrastructure/repository/mongo/customer"
	"meliarqsoft2/internal/infrastructure/repository/mongo/seller"
	"meliarqsoft2/internal/infrastructure/repository/rabbitmq"
	"meliarqsoft2/internal/infrastructure/repository/rabbitmq/notification"
	"meliarqsoft2/internal/infrastructure/service"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port, err := strconv.Atoi(os.Getenv("USERS_APP_PORT"))
	if err != nil {
		panic(err.Error())
	}

	sellerClient := mongo.NewFactory().InitMongoDB(os.Getenv("MONGODB_SELLERS_URI"))
	customerClient := mongo.NewFactory().InitMongoDB(os.Getenv("MONGODB_CUSTOMERS_URI"))

	sellerRepository := seller.NewMongoRepository(sellerClient)
	customerRepository := customer.NewMongoRepository(customerClient)

	// RabbitMQ

	clientMQ := rabbitmq.NewFactory().InitRabbitMQ()

	notificationRepository := notification.NewRabbitMQRepository(clientMQ)

	// seller
	registerSellerEvent := action.NewRegisterSellerEvent(sellerRepository)
	updateSellerEvent := action.NewUpdateSellerEvent(sellerRepository)
	findSellerEvent := query.NewFindSellerEvent(sellerRepository)
	findSellerByIdEvent := query.NewFindSellerByIdEvent(sellerRepository)
	deleteProdsService := service.NewHttpSyncDeleteProductsBySeller(os.Getenv("PRODUCT_URL"))
	unregisterSellerEvent := action.NewUnregisterSellerEvent(sellerRepository, deleteProdsService)

	// user
	findUserById := query.NewFindCustomerByIdEvent(customerRepository)
	registerCustomerEvent := action.NewRegisterCustomerEvent(customerRepository)
	updateCustomerEvent := action.NewUpdateCustomerEvent(customerRepository)
	findCustomerEvent := query.NewFindCustomerEvent(customerRepository)
	unregisterCustomerEvent := action.NewUnregisterCustomerEvent(customerRepository)

	// Notification
	sendNotification := action.NewSendNotificationEvent(notificationRepository)

	newAPI := gin.NewMeliAPI(
		port,
		api.NewEvents(
			registerSellerEvent,
			updateSellerEvent,
			unregisterSellerEvent,
			findSellerEvent,
			findSellerByIdEvent,

			registerCustomerEvent,
			updateCustomerEvent,
			unregisterCustomerEvent,
			findCustomerEvent,
			findUserById,

			sendNotification,
		),
	)

	if os.Getenv("ENVIRONMENT") == "test" {
		seed(sellerRepository, customerRepository)
	}

	err = newAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}

func seed(sellerRepository model.ISellerRepository, customerRepository model.ICustomerRepository) {
	id, _ := uuid.Parse("9d032089-9223-4c98-9196-f97c6f792473")

	email, _ := model.NewEmail("vendedor@prueba.com")
	err := sellerRepository.DeleteAll()
	if err != nil && err != mongo2.ErrNoDocuments {
		panic("error: " + err.Error())
	}

	_, err = sellerRepository.Create(&model.Seller{
		ID:           id,
		BusinessName: "VENTA PRUEBAS",
		Email:        email,
	})
	if err != nil {
		panic("failed in user service setup: " + err.Error())
	}

	email, _ = model.NewEmail("cliente@prueba.com")
	err = customerRepository.DeleteAll()
	if err != nil && err != mongo2.ErrNoDocuments {
		panic("error: " + err.Error())
	}

	_, err = customerRepository.Create(&model.Customer{
		ID:      id,
		Name:    "prueba",
		Surname: "prueba",
		Email:   email,
	})
	if err != nil {
		panic("failed in user service setup: " + err.Error())
	}
}
