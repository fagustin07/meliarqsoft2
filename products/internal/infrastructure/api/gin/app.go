package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"meliarqsoft2/docs"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin/handler"
	"net/http"
	"strconv"
)

type MeliGinApp struct {
	port   int
	events *api.Events
}

func NewMeliAPI(port int, events *api.Events) *MeliGinApp {
	return &MeliGinApp{
		port:   port,
		events: events,
	}
}

func (app MeliGinApp) Run() error {

	route := gin.Default()

	// swagger config
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "MELI - Product Service"
	docs.SwaggerInfo.Description = "Proyecto de Arquitectura de Software 2, UNQ 2023s1"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + strconv.Itoa(app.port)
	docs.SwaggerInfo.BasePath = "/api/v1"

	route.GET("/api/v1/heartbeat", heartbeat)

	productRoute := route.Group("/api/v1/products")

	productRoute.POST("", handler.NewGinCreateProduct(app.events.CreateProductEvent).Execute)
	productRoute.GET("", handler.NewGinFindProduct(app.events.FindProductEvent).Execute)
	productRoute.PUT("/:id", handler.NewGinUpdateProduct(app.events.UpdateProductEvent).Execute)
	productRoute.GET("/:id", handler.NewGinFindProdById(app.events.FindProductById).Execute)
	productRoute.DELETE("/:id", handler.NewGinDeleteProduct(app.events.DeleteProductEvent).Execute)
	productRoute.GET("/prices", handler.NewGinFilterProduct(app.events.FilterProductEvent).Execute)

	productRoute.DELETE("/sellers/:id", handler.NewGinDeleteProductsBySeller(app.events.DeleteProductsBySeller).Execute)
	productRoute.POST("/restore", handler.NewGinRestoreProducts(app.events.RestoreProductsEvent).Execute)

	return route.Run(fmt.Sprintf(":%d", app.port))
}

func heartbeat(context *gin.Context) {
	context.JSON(http.StatusOK, "PRODUCTS SERVICE IS WORKING!")
}
