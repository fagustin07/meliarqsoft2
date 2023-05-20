package main

import (
	"github.com/joho/godotenv"
	"log"
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin"
	"meliarqsoft2/internal/infrastructure/repository/mongo"
	"meliarqsoft2/internal/infrastructure/repository/mongo/seller"
	"meliarqsoft2/internal/infrastructure/repository/mongo/user"
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
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err.Error())
	}

	client := mongo.NewFactory().InitMongoDB()

	sellerRepository := seller.NewMongoRepository(client)
	userRepository := user.NewMongoRepository(client)

	// RabbitMQ

	clientMQ := rabbitmq.NewFactory().InitRabbitMQ()

	notificationRepository := notification.NewRabbitMQRepository(clientMQ)

	// seller
	registerSellerEvent := action.NewRegisterSellerEvent(sellerRepository)
	updateSellerEvent := action.NewUpdateSellerEvent(sellerRepository)
	existSeller := query.NewExistSellerCommand(sellerRepository)
	unregisterSellerEvent := action.NewUnregisterSellerEvent(sellerRepository, existSeller)
	findSellerEvent := query.NewFindSellerEvent(sellerRepository)

	// user
	existUser := query.NewExistUserCommand(userRepository)
	registerUserEvent := action.NewRegisterUserEvent(userRepository)
	updateUserEvent := action.NewUpdateUserEvent(userRepository)
	findUserEvent := query.NewFindUserEvent(userRepository)
	unregisterUserEvent := action.NewUnregisterUserEvent(userRepository)

	// Notification
	sendNotification := action.NewSendNotificationEvent(notificationRepository)

	newAPI := gin.NewMeliAPI(
		port,
		api.NewEvents(
			registerSellerEvent,
			updateSellerEvent,
			unregisterSellerEvent,
			findSellerEvent,
			existSeller,

			registerUserEvent,
			updateUserEvent,
			unregisterUserEvent,
			findUserEvent,
			existUser,

			sendNotification,
		),
	)

	err = newAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}
