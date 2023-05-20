package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"meliarqsoft2/docs"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin/handler"
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

	userRoute := basePath.Group("/users")
	userRoute.POST("", handler.NewGinUserRegister(app.events.RegisterUserEvent, app.events.SendNotificationEvent).Execute)
	userRoute.GET("", handler.NewGinFindUser(app.events.FindUserEvent).Execute)
	userRoute.PUT("/:id", handler.NewGinUpdateUser(app.events.UpdateUserEvent).Execute)
	userRoute.DELETE("/:id", handler.NewGinUnregisterUser(app.events.UnregisterUserEvent).Execute)

	sellerRoute := basePath.Group("/sellers")
	sellerRoute.POST("", handler.NewGinRegisterSeller(app.events.RegisterSellerEvent, app.events.SendNotificationEvent).Execute)
	sellerRoute.GET("", handler.NewGinFindSeller(app.events.FindSellerEvent).Execute)
	sellerRoute.PUT("/:id", handler.NewGinUpdateSeller(app.events.UpdateSellerEvent).Execute)
	sellerRoute.DELETE("/:id", handler.NewGinUnregisterSeller(app.events.UnregisterSellerEvent).Execute)

	return route.Run(fmt.Sprintf(":%d", app.port))
}
