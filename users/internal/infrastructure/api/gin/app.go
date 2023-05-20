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
	docs.SwaggerInfo.Title = "MELI - User Microservice"
	docs.SwaggerInfo.Description = "Proyecto de Arquitectura de Software 2, UNQ 2023s1"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + strconv.Itoa(app.port)
	docs.SwaggerInfo.BasePath = "/api/v1"

	basePath := route.Group("/api/v1")

	basePath.GET("/heartbeat", heartbeat)
	customerRoute := basePath.Group("/customers")
	customerRoute.POST("", handler.NewGinCustomerRegister(app.events.RegisterCustomerEvent, app.events.SendNotificationEvent).Execute)
	customerRoute.GET("", handler.NewGinFindCustomer(app.events.FindCustomerEvent).Execute)
	customerRoute.PUT("/:id", handler.NewGinUpdateCustomer(app.events.UpdateCustomerEvent).Execute)
	customerRoute.DELETE("/:id", handler.NewGinUnregisterCustomer(app.events.UnregisterCustomerEvent).Execute)

	sellerRoute := basePath.Group("/sellers")
	sellerRoute.POST("", handler.NewGinRegisterSeller(app.events.RegisterSellerEvent, app.events.SendNotificationEvent).Execute)
	sellerRoute.GET("", handler.NewGinFindSeller(app.events.FindSellerEvent).Execute)
	sellerRoute.PUT("/:id", handler.NewGinUpdateSeller(app.events.UpdateSellerEvent).Execute)
	sellerRoute.DELETE("/:id", handler.NewGinUnregisterSeller(app.events.UnregisterSellerEvent).Execute)

	return route.Run(fmt.Sprintf(":%d", app.port))
}

func heartbeat(context *gin.Context) {
	context.JSON(http.StatusOK, "USER SERVICE IS WORKING!")
}
