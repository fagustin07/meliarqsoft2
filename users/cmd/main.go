package main

import (
	"github.com/joho/godotenv"
	"log"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin"
	"meliarqsoft2/internal/infrastructure/repository/mongo"
	"meliarqsoft2/internal/infrastructure/repository/mongo/customer"
	"meliarqsoft2/internal/infrastructure/repository/mongo/seller"
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
	unregisterSellerEvent := action.NewUnregisterSellerEvent(sellerRepository)
	findSellerEvent := query.NewFindSellerEvent(sellerRepository)
	findSellerByIdEvent := query.NewFindSellerByIdEvent(sellerRepository)

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

	err = newAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}
