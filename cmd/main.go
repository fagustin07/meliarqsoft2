package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"meliarqsoft2/config/factory"
	"meliarqsoft2/docs"
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
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	basePath := r.Group("/api/v1")
	productRoute := basePath.Group("/products")

	configProductRoutes(productRoute, productHandler)

	docs.SwaggerInfo.Title = "MELI - Arquitectura Hexagonal"
	docs.SwaggerInfo.Description = "Proyecto de Arquitectura de Software 2, UNQ 2023s1"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	otherErr := r.Run()
	if otherErr != nil {
		return
	}
}

func configProductRoutes(route *gin.RouterGroup, ginHandler *handler.ProductGinHandler) {
	route.POST("", ginHandler.Create)
	route.PUT("/:id", ginHandler.Update)
	route.DELETE("/:id", ginHandler.Delete)
	route.GET("", ginHandler.Find)
	route.GET("/prices", ginHandler.Filter)
}
