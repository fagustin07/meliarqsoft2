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
	handler4 "meliarqsoft2/internal/purchase/infrastructure/handler"
	handler2 "meliarqsoft2/internal/seller/infrastructure/handler"
	handler3 "meliarqsoft2/internal/user/infrastructure/handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newFactory := factory.NewFactory()

	newFactory.InitMongoDB()

	_, sellerHandler := newFactory.BuildSellerHandler()
	_, userHandler := newFactory.BuildUserHandler()
	_, productHandler := newFactory.BuildProductHandler()
	_, purchaseHandler := newFactory.BuildPurchaseHandler()

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	basePath := r.Group("/api/v1")

	userRoute := basePath.Group("/users")
	configUserRoutes(userRoute, userHandler)

	sellerRoute := basePath.Group("/sellers")
	configSellerRoutes(sellerRoute, sellerHandler)

	productRoute := basePath.Group("/products")
	configProductRoutes(productRoute, productHandler, purchaseHandler)

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

func configProductRoutes(route *gin.RouterGroup, prodHandler *handler.ProductGinHandler, purchaseHandler *handler4.PurchaseGinHandler) {
	route.POST("", prodHandler.Create)
	route.PUT("/:id", prodHandler.Update)
	route.DELETE("/:id", prodHandler.Delete)
	route.GET("", prodHandler.Find)
	route.GET("/prices", prodHandler.Filter)

	route.POST("/purchases", purchaseHandler.Create)
	route.GET("/:id/purchases", purchaseHandler.Find)
}

func configUserRoutes(route *gin.RouterGroup, ginHandler *handler3.UserGinHandler) {
	route.POST("", ginHandler.Create)
	route.GET("", ginHandler.Find)
	route.PUT("/:id", ginHandler.Update)
	route.DELETE("/:id", ginHandler.Delete)
}

func configSellerRoutes(route *gin.RouterGroup, ginHandler *handler2.SellerGinHandler) {
	route.POST("", ginHandler.Create)
	route.PUT("/:id", ginHandler.Update)
	route.DELETE("/:id", ginHandler.Delete)
	route.GET("", ginHandler.Find)
}
