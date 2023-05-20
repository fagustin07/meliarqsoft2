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
)

type MeliGinApp struct {
	port   int
	events *api.PurchaseEvents
}

func NewMeliAPI(port int, events *api.PurchaseEvents) *MeliGinApp {
	return &MeliGinApp{
		port:   port,
		events: events,
	}
}

func (app MeliGinApp) Run() error {

	route := gin.Default()

	// swagger config
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "MELI - Purchase Microservice"
	docs.SwaggerInfo.Description = "Proyecto de Arquitectura de Software 2, UNQ 2023s1"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", app.port)
	docs.SwaggerInfo.BasePath = "/api/v1"

	basePath := route.Group("/api/v1")
	basePath.GET("/heartbeat", heartbeat)

	productRoute := basePath.Group("/purchases")
	productRoute.DELETE("/:id", handler.NewGinDeleteFromProduct(app.events.UndoPurchasesFromProductEvent).Execute)
	productRoute.POST("/purchases", handler.NewGinMakePurchase(app.events.MakePurchaseEvent, app.events.SendNotificationEvent).Execute)
	productRoute.GET("/products/:id", handler.NewGinFindPurchases(app.events.FindPurchasesFromProductEvent).Execute)

	return route.Run(fmt.Sprintf(":%d", app.port))
}

func heartbeat(context *gin.Context) {
	context.JSON(http.StatusOK, "PURCHASES SERVICE IS WORKING!")
}
