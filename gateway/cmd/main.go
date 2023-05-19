package main

import (
	"github.com/joho/godotenv"
	"log"
	"meliarqsoft2/internal/infrastructure/api/gin"
	"meliarqsoft2/internal/infrastructure/services"
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

	newAPI := gin.NewMeliAPIGateway(
		port,
		services.CustomerHttpSyncService{BasePath: os.Getenv("USER_URL")},
		services.SellerHttpSyncService{BasePath: os.Getenv("USER_URL")},
		services.ProductHttpSyncService{BasePath: os.Getenv("PRODUCT_URL")},
		services.FindPurchaseHttpSyncService{BasePath: os.Getenv("PURCHASE_URL")},
	)

	err = newAPI.Run()

	if err != nil {
		log.Fatal("failed running app")
	}
}
