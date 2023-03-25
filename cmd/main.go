package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"meliarqsoft2/config/factory"
	"meliarqsoft2/internal/product/infrastructure/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newFactory := factory.NewFactory()

	newFactory.InitMongoDB()

	productHandler := newFactory.BuildProductHandler()

	r := gin.Default()
	basePath := r.Group("/api/v1")
	productRoute := basePath.Group("/product")

	configProductRoutes(productRoute, productHandler)
	r.GET("/ping", func(c *gin.Context) {
		newUUID, _ := uuid.NewUUID()
		c.JSON(200, gin.H{
			"uuid": newUUID,
		})
	})

	otherErr := r.Run()
	if otherErr != nil {
		return
	}
}

func configProductRoutes(route *gin.RouterGroup, ginHandler *handler.ProductGinHandler) {
	route.POST("", ginHandler.Create)
	route.GET("", ginHandler.Find)
}
